/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package unstructured

import (
	gojson "encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
)

func getNestedField(obj map[string]interface{}, fields ...string) interface{} {
	var val interface{} = obj
	for _, field := range fields {
		if _, ok := val.(map[string]interface{}); !ok {
			return nil
		}
		val = val.(map[string]interface{})[field]
	}
	return val
}

func getNestedString(obj map[string]interface{}, fields ...string) string {
	if str, ok := getNestedField(obj, fields...).(string); ok {
		return str
	}
	return ""
}

func getNestedInt64(obj map[string]interface{}, fields ...string) int64 {
	if str, ok := getNestedField(obj, fields...).(int64); ok {
		return str
	}
	return 0
}

func getNestedInt64Pointer(obj map[string]interface{}, fields ...string) *int64 {
	nested := getNestedField(obj, fields...)
	switch n := nested.(type) {
	case int64:
		return &n
	case *int64:
		return n
	default:
		return nil
	}
}

func getNestedSlice(obj map[string]interface{}, fields ...string) []string {
	if m, ok := getNestedField(obj, fields...).([]interface{}); ok {
		strSlice := make([]string, 0, len(m))
		for _, v := range m {
			if str, ok := v.(string); ok {
				strSlice = append(strSlice, str)
			}
		}
		return strSlice
	}
	return nil
}

func getNestedMap(obj map[string]interface{}, fields ...string) map[string]string {
	if m, ok := getNestedField(obj, fields...).(map[string]interface{}); ok {
		strMap := make(map[string]string, len(m))
		for k, v := range m {
			if str, ok := v.(string); ok {
				strMap[k] = str
			}
		}
		return strMap
	}
	return nil
}

func setNestedField(obj map[string]interface{}, value interface{}, fields ...string) {
	m := obj
	if len(fields) > 1 {
		for _, field := range fields[0 : len(fields)-1] {
			if _, ok := m[field].(map[string]interface{}); !ok {
				m[field] = make(map[string]interface{})
			}
			m = m[field].(map[string]interface{})
		}
	}
	m[fields[len(fields)-1]] = value
}

func setNestedSlice(obj map[string]interface{}, value []string, fields ...string) {
	m := make([]interface{}, 0, len(value))
	for _, v := range value {
		m = append(m, v)
	}
	setNestedField(obj, m, fields...)
}

func setNestedMap(obj map[string]interface{}, value map[string]string, fields ...string) {
	m := make(map[string]interface{}, len(value))
	for k, v := range value {
		m[k] = v
	}
	setNestedField(obj, m, fields...)
}

func extractOwnerReference(src interface{}) metav1.OwnerReference {
	v := src.(map[string]interface{})
	// though this field is a *bool, but when decoded from JSON, it's
	// unmarshalled as bool.
	var controllerPtr *bool
	controller, ok := (getNestedField(v, "controller")).(bool)
	if !ok {
		controllerPtr = nil
	} else {
		controllerCopy := controller
		controllerPtr = &controllerCopy
	}
	var blockOwnerDeletionPtr *bool
	blockOwnerDeletion, ok := (getNestedField(v, "blockOwnerDeletion")).(bool)
	if !ok {
		blockOwnerDeletionPtr = nil
	} else {
		blockOwnerDeletionCopy := blockOwnerDeletion
		blockOwnerDeletionPtr = &blockOwnerDeletionCopy
	}
	return metav1.OwnerReference{
		Kind:               getNestedString(v, "kind"),
		Name:               getNestedString(v, "name"),
		APIVersion:         getNestedString(v, "apiVersion"),
		UID:                (types.UID)(getNestedString(v, "uid")),
		Controller:         controllerPtr,
		BlockOwnerDeletion: blockOwnerDeletionPtr,
	}
}

func setOwnerReference(src metav1.OwnerReference) map[string]interface{} {
	ret := make(map[string]interface{})
	setNestedField(ret, src.Kind, "kind")
	setNestedField(ret, src.Name, "name")
	setNestedField(ret, src.APIVersion, "apiVersion")
	setNestedField(ret, string(src.UID), "uid")
	// json.Unmarshal() extracts boolean json fields as bool, not as *bool and hence extractOwnerReference()
	// expects bool or a missing field, not *bool. So if pointer is nil, fields are omitted from the ret object.
	// If pointer is non-nil, they are set to the referenced value.
	if src.Controller != nil {
		setNestedField(ret, *src.Controller, "controller")
	}
	if src.BlockOwnerDeletion != nil {
		setNestedField(ret, *src.BlockOwnerDeletion, "blockOwnerDeletion")
	}
	return ret
}

func getOwnerReferences(object map[string]interface{}) ([]map[string]interface{}, error) {
	field := getNestedField(object, "metadata", "ownerReferences")
	if field == nil {
		return nil, fmt.Errorf("cannot find field metadata.ownerReferences in %v", object)
	}
	ownerReferences, ok := field.([]map[string]interface{})
	if ok {
		return ownerReferences, nil
	}
	// TODO: This is hacky...
	interfaces, ok := field.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expect metadata.ownerReferences to be a slice in %#v", object)
	}
	ownerReferences = make([]map[string]interface{}, 0, len(interfaces))
	for i := 0; i < len(interfaces); i++ {
		r, ok := interfaces[i].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expect element metadata.ownerReferences to be a map[string]interface{} in %#v", object)
		}
		ownerReferences = append(ownerReferences, r)
	}
	return ownerReferences, nil
}

var converter = unstructured.NewConverter(false)

// UnstructuredJSONScheme is capable of converting JSON data into the Unstructured
// type, which can be used for generic access to objects without a predefined scheme.
// TODO: move into serializer/json.
var UnstructuredJSONScheme runtime.Codec = unstructuredJSONScheme{}

type unstructuredJSONScheme struct{}

func (s unstructuredJSONScheme) Decode(data []byte, _ *schema.GroupVersionKind, obj runtime.Object) (runtime.Object, *schema.GroupVersionKind, error) {
	var err error
	if obj != nil {
		err = s.decodeInto(data, obj)
	} else {
		obj, err = s.decode(data)
	}

	if err != nil {
		return nil, nil, err
	}

	gvk := obj.GetObjectKind().GroupVersionKind()
	if len(gvk.Kind) == 0 {
		return nil, &gvk, runtime.NewMissingKindErr(string(data))
	}

	return obj, &gvk, nil
}

func (unstructuredJSONScheme) Encode(obj runtime.Object, w io.Writer) error {
	switch t := obj.(type) {
	case *Unstructured:
		return json.NewEncoder(w).Encode(t.Object)
	case *UnstructuredList:
		items := make([]map[string]interface{}, 0, len(t.Items))
		for _, i := range t.Items {
			items = append(items, i.Object)
		}
		listObj := make(map[string]interface{}, len(t.Object)+1)
		for k, v := range t.Object { // Make a shallow copy
			listObj[k] = v
		}
		listObj["items"] = items
		return json.NewEncoder(w).Encode(listObj)
	case *runtime.Unknown:
		// TODO: Unstructured needs to deal with ContentType.
		_, err := w.Write(t.Raw)
		return err
	default:
		return json.NewEncoder(w).Encode(t)
	}
}

func (s unstructuredJSONScheme) decode(data []byte) (runtime.Object, error) {
	type detector struct {
		Items gojson.RawMessage
	}
	var det detector
	if err := json.Unmarshal(data, &det); err != nil {
		return nil, err
	}

	if det.Items != nil {
		list := &UnstructuredList{}
		err := s.decodeToList(data, list)
		return list, err
	}

	// No Items field, so it wasn't a list.
	unstruct := &Unstructured{}
	err := s.decodeToUnstructured(data, unstruct)
	return unstruct, err
}

func (s unstructuredJSONScheme) decodeInto(data []byte, obj runtime.Object) error {
	switch x := obj.(type) {
	case *Unstructured:
		return s.decodeToUnstructured(data, x)
	case *UnstructuredList:
		return s.decodeToList(data, x)
	case *runtime.VersionedObjects:
		o, err := s.decode(data)
		if err == nil {
			x.Objects = []runtime.Object{o}
		}
		return err
	default:
		return json.Unmarshal(data, x)
	}
}

func (unstructuredJSONScheme) decodeToUnstructured(data []byte, unstruct *Unstructured) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	unstruct.Object = m

	return nil
}

func (s unstructuredJSONScheme) decodeToList(data []byte, list *UnstructuredList) error {
	type decodeList struct {
		Items []gojson.RawMessage
	}

	var dList decodeList
	if err := json.Unmarshal(data, &dList); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &list.Object); err != nil {
		return err
	}

	// For typed lists, e.g., a PodList, API server doesn't set each item's
	// APIVersion and Kind. We need to set it.
	listAPIVersion := list.GetAPIVersion()
	listKind := list.GetKind()
	itemKind := strings.TrimSuffix(listKind, "List")

	delete(list.Object, "items")
	list.Items = nil
	for _, i := range dList.Items {
		unstruct := &Unstructured{}
		if err := s.decodeToUnstructured([]byte(i), unstruct); err != nil {
			return err
		}
		// This is hacky. Set the item's Kind and APIVersion to those inferred
		// from the List.
		if len(unstruct.GetKind()) == 0 && len(unstruct.GetAPIVersion()) == 0 {
			unstruct.SetKind(itemKind)
			unstruct.SetAPIVersion(listAPIVersion)
		}
		list.Items = append(list.Items, *unstruct)
	}
	return nil
}

// UnstructuredObjectConverter is an ObjectConverter for use with
// Unstructured objects. Since it has no schema or type information,
// it will only succeed for no-op conversions. This is provided as a
// sane implementation for APIs that require an object converter.
type UnstructuredObjectConverter struct{}

func (UnstructuredObjectConverter) Convert(in, out, context interface{}) error {
	unstructIn, ok := in.(*Unstructured)
	if !ok {
		return fmt.Errorf("input type %T in not valid for unstructured conversion", in)
	}

	unstructOut, ok := out.(*Unstructured)
	if !ok {
		return fmt.Errorf("output type %T in not valid for unstructured conversion", out)
	}

	// maybe deep copy the map? It is documented in the
	// ObjectConverter interface that this function is not
	// guaranteeed to not mutate the input. Or maybe set the input
	// object to nil.
	unstructOut.Object = unstructIn.Object
	return nil
}

func (UnstructuredObjectConverter) ConvertToVersion(in runtime.Object, target runtime.GroupVersioner) (runtime.Object, error) {
	if kind := in.GetObjectKind().GroupVersionKind(); !kind.Empty() {
		gvk, ok := target.KindForGroupVersionKinds([]schema.GroupVersionKind{kind})
		if !ok {
			// TODO: should this be a typed error?
			return nil, fmt.Errorf("%v is unstructured and is not suitable for converting to %q", kind, target)
		}
		in.GetObjectKind().SetGroupVersionKind(gvk)
	}
	return in, nil
}

func (UnstructuredObjectConverter) ConvertFieldLabel(version, kind, label, value string) (string, string, error) {
	return "", "", errors.New("unstructured cannot convert field labels")
}
