// +build !ignore_autogenerated

/*
Copyright 2019 The Knative Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	duckv1alpha1 "knative.dev/eventing/pkg/apis/duck/v1alpha1"
	eventingv1alpha1 "knative.dev/eventing/pkg/apis/eventing/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Channel) DeepCopyInto(out *Channel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Channel.
func (in *Channel) DeepCopy() *Channel {
	if in == nil {
		return nil
	}
	out := new(Channel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Channel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelList) DeepCopyInto(out *ChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Channel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelList.
func (in *ChannelList) DeepCopy() *ChannelList {
	if in == nil {
		return nil
	}
	out := new(ChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelSpec) DeepCopyInto(out *ChannelSpec) {
	*out = *in
	if in.ChannelTemplate != nil {
		in, out := &in.ChannelTemplate, &out.ChannelTemplate
		*out = new(duckv1alpha1.ChannelTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Subscribable != nil {
		in, out := &in.Subscribable, &out.Subscribable
		*out = new(duckv1alpha1.Subscribable)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelSpec.
func (in *ChannelSpec) DeepCopy() *ChannelSpec {
	if in == nil {
		return nil
	}
	out := new(ChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelStatus) DeepCopyInto(out *ChannelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	in.SubscribableTypeStatus.DeepCopyInto(&out.SubscribableTypeStatus)
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelStatus.
func (in *ChannelStatus) DeepCopy() *ChannelStatus {
	if in == nil {
		return nil
	}
	out := new(ChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Choice) DeepCopyInto(out *Choice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Choice.
func (in *Choice) DeepCopy() *Choice {
	if in == nil {
		return nil
	}
	out := new(Choice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Choice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChoiceCase) DeepCopyInto(out *ChoiceCase) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(eventingv1alpha1.SubscriberSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Subscriber.DeepCopyInto(&out.Subscriber)
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChoiceCase.
func (in *ChoiceCase) DeepCopy() *ChoiceCase {
	if in == nil {
		return nil
	}
	out := new(ChoiceCase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChoiceCaseStatus) DeepCopyInto(out *ChoiceCaseStatus) {
	*out = *in
	in.FilterSubscriptionStatus.DeepCopyInto(&out.FilterSubscriptionStatus)
	in.FilterChannelStatus.DeepCopyInto(&out.FilterChannelStatus)
	in.SubscriptionStatus.DeepCopyInto(&out.SubscriptionStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChoiceCaseStatus.
func (in *ChoiceCaseStatus) DeepCopy() *ChoiceCaseStatus {
	if in == nil {
		return nil
	}
	out := new(ChoiceCaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChoiceChannelStatus) DeepCopyInto(out *ChoiceChannelStatus) {
	*out = *in
	out.Channel = in.Channel
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChoiceChannelStatus.
func (in *ChoiceChannelStatus) DeepCopy() *ChoiceChannelStatus {
	if in == nil {
		return nil
	}
	out := new(ChoiceChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChoiceList) DeepCopyInto(out *ChoiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Choice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChoiceList.
func (in *ChoiceList) DeepCopy() *ChoiceList {
	if in == nil {
		return nil
	}
	out := new(ChoiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChoiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChoiceSpec) DeepCopyInto(out *ChoiceSpec) {
	*out = *in
	if in.Cases != nil {
		in, out := &in.Cases, &out.Cases
		*out = make([]ChoiceCase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChannelTemplate != nil {
		in, out := &in.ChannelTemplate, &out.ChannelTemplate
		*out = new(duckv1alpha1.ChannelTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChoiceSpec.
func (in *ChoiceSpec) DeepCopy() *ChoiceSpec {
	if in == nil {
		return nil
	}
	out := new(ChoiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChoiceStatus) DeepCopyInto(out *ChoiceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.IngressChannelStatus.DeepCopyInto(&out.IngressChannelStatus)
	if in.CaseStatuses != nil {
		in, out := &in.CaseStatuses, &out.CaseStatuses
		*out = make([]ChoiceCaseStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChoiceStatus.
func (in *ChoiceStatus) DeepCopy() *ChoiceStatus {
	if in == nil {
		return nil
	}
	out := new(ChoiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChoiceSubscriptionStatus) DeepCopyInto(out *ChoiceSubscriptionStatus) {
	*out = *in
	out.Subscription = in.Subscription
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChoiceSubscriptionStatus.
func (in *ChoiceSubscriptionStatus) DeepCopy() *ChoiceSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(ChoiceSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannel) DeepCopyInto(out *InMemoryChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannel.
func (in *InMemoryChannel) DeepCopy() *InMemoryChannel {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannelList) DeepCopyInto(out *InMemoryChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InMemoryChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannelList.
func (in *InMemoryChannelList) DeepCopy() *InMemoryChannelList {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannelSpec) DeepCopyInto(out *InMemoryChannelSpec) {
	*out = *in
	if in.Subscribable != nil {
		in, out := &in.Subscribable, &out.Subscribable
		*out = new(duckv1alpha1.Subscribable)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannelSpec.
func (in *InMemoryChannelSpec) DeepCopy() *InMemoryChannelSpec {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannelStatus) DeepCopyInto(out *InMemoryChannelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	in.SubscribableTypeStatus.DeepCopyInto(&out.SubscribableTypeStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannelStatus.
func (in *InMemoryChannelStatus) DeepCopy() *InMemoryChannelStatus {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sequence) DeepCopyInto(out *Sequence) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sequence.
func (in *Sequence) DeepCopy() *Sequence {
	if in == nil {
		return nil
	}
	out := new(Sequence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sequence) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceChannelStatus) DeepCopyInto(out *SequenceChannelStatus) {
	*out = *in
	out.Channel = in.Channel
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceChannelStatus.
func (in *SequenceChannelStatus) DeepCopy() *SequenceChannelStatus {
	if in == nil {
		return nil
	}
	out := new(SequenceChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceList) DeepCopyInto(out *SequenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sequence, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceList.
func (in *SequenceList) DeepCopy() *SequenceList {
	if in == nil {
		return nil
	}
	out := new(SequenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SequenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceSpec) DeepCopyInto(out *SequenceSpec) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]eventingv1alpha1.SubscriberSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChannelTemplate != nil {
		in, out := &in.ChannelTemplate, &out.ChannelTemplate
		*out = new(duckv1alpha1.ChannelTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceSpec.
func (in *SequenceSpec) DeepCopy() *SequenceSpec {
	if in == nil {
		return nil
	}
	out := new(SequenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceStatus) DeepCopyInto(out *SequenceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.SubscriptionStatuses != nil {
		in, out := &in.SubscriptionStatuses, &out.SubscriptionStatuses
		*out = make([]SequenceSubscriptionStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChannelStatuses != nil {
		in, out := &in.ChannelStatuses, &out.ChannelStatuses
		*out = make([]SequenceChannelStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceStatus.
func (in *SequenceStatus) DeepCopy() *SequenceStatus {
	if in == nil {
		return nil
	}
	out := new(SequenceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceSubscriptionStatus) DeepCopyInto(out *SequenceSubscriptionStatus) {
	*out = *in
	out.Subscription = in.Subscription
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceSubscriptionStatus.
func (in *SequenceSubscriptionStatus) DeepCopy() *SequenceSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SequenceSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}
