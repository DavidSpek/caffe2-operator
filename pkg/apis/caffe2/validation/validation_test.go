// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import (
	"testing"

	caffe2v1 "github.com/kubeflow/caffe2-operator/pkg/apis/caffe2/v1alpha1"
	"github.com/kubeflow/caffe2-operator/pkg/util"

	"github.com/gogo/protobuf/proto"
	"k8s.io/api/core/v1"
)

func TestValidate(t *testing.T) {
	type testCase struct {
		in             *caffe2v1.Caffe2JobSpec
		expectingError bool
	}

	testCases := []testCase{
		{
			in: &caffe2v1.Caffe2JobSpec{
				ReplicaSpecs: &caffe2v1.Caffe2ReplicaSpec{
					Template: &v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								{
									Name: "tensorflow",
								},
							},
						},
					},
					Replicas: proto.Int32(1),
				},
			},
			expectingError: false,
		},
		{
			in: &caffe2v1.Caffe2JobSpec{
				ReplicaSpecs: &caffe2v1.Caffe2ReplicaSpec{
					Template: &v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								{
									Name: "tensorflow",
								},
							},
						},
					},
					Replicas: proto.Int32(1),
				},
				TerminationPolicy: &caffe2v1.TerminationPolicySpec{
					Chief: &caffe2v1.ChiefSpec{
						ReplicaName:  "WORKER",
						ReplicaIndex: 0,
					},
				},
			},
			expectingError: false,
		},
	}

	for _, c := range testCases {
		job := &caffe2v1.Caffe2Job{
			Spec: *c.in,
		}
		caffe2v1.SetObjectDefaults_Caffe2Job(job)
		if err := ValidateCaffe2JobSpec(&job.Spec); (err != nil) != c.expectingError {
			t.Logf("job.Spec: %v", util.Pformat(job.Spec))
			t.Errorf("unexpected validation result: %v", err)
		}
	}
}
