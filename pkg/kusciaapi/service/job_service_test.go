// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:dupl
package service

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	consts "github.com/secretflow/kuscia/pkg/web/constants"
	"github.com/secretflow/kuscia/proto/api/v1alpha1/errorcode"
	"github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapi"
)

func TestCreateJob(t *testing.T) {
	res := kusciaAPIJS.CreateJob(context.Background(), &kusciaapi.CreateJobRequest{
		JobId:     kusciaAPIJS.jobID,
		Initiator: "alice",
		Tasks:     kusciaAPIJS.tasks,
	})
	assert.Equal(t, res.Data.JobId, kusciaAPIJS.jobID)
}

func TestQueryJob(t *testing.T) {
	queryJobResponse := kusciaAPIJS.QueryJob(context.Background(), &kusciaapi.QueryJobRequest{
		JobId: kusciaAPIJS.jobID,
	})
	assert.Equal(t, queryJobResponse.Data.JobId, kusciaAPIJS.jobID)
	assert.Equal(t, len(queryJobResponse.Data.Tasks), len(kusciaAPIJS.tasks))
}

func TestSuspendJob(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.AuthRole, consts.AuthRoleMaster)
	ctx = context.WithValue(ctx, consts.SourceDomainKey, "alice")
	kusciaAPIJS.CreateJob(ctx, &kusciaapi.CreateJobRequest{
		JobId:     kusciaAPIJS.jobID,
		Initiator: "alice",
		Tasks:     kusciaAPIJS.tasks,
	})
	res := kusciaAPIJS.SuspendJob(ctx, &kusciaapi.SuspendJobRequest{
		JobId: kusciaAPIJS.jobID,
	})
	assert.Equal(t, res.Status.Code == 0, false)
	kj, err := kusciaClient.KusciaV1alpha1().KusciaJobs("cross-domain").Get(ctx, kusciaAPIJS.jobID, metav1.GetOptions{})
	if err != nil {
		t.Error(err.Error())
	}
	kj.Status.Phase = v1alpha1.KusciaJobRunning
	kj, err = kusciaClient.KusciaV1alpha1().KusciaJobs("cross-domain").UpdateStatus(ctx, kj, metav1.UpdateOptions{})
	if err != nil {
		t.Error(err.Error())
	}
	res = kusciaAPIJS.SuspendJob(ctx, &kusciaapi.SuspendJobRequest{
		JobId: kusciaAPIJS.jobID,
	})
	assert.Equal(t, res.Data.JobId, kusciaAPIJS.jobID)
}

func TestRestartJob(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.AuthRole, consts.AuthRoleMaster)
	ctx = context.WithValue(ctx, consts.SourceDomainKey, "alice")
	res := kusciaAPIJS.RestartJob(ctx, &kusciaapi.RestartJobRequest{
		JobId: kusciaAPIJS.jobID,
	})
	assert.Equal(t, res.Status.Code == 0, false)
	kj, err := kusciaClient.KusciaV1alpha1().KusciaJobs("cross-domain").Get(ctx, kusciaAPIJS.jobID, metav1.GetOptions{})
	if err != nil {
		t.Error(err.Error())
	}
	kj.Status.Phase = v1alpha1.KusciaJobFailed
	kj, err = kusciaClient.KusciaV1alpha1().KusciaJobs("cross-domain").UpdateStatus(ctx, kj, metav1.UpdateOptions{})
	if err != nil {
		t.Error(err.Error())
	}
	res = kusciaAPIJS.RestartJob(ctx, &kusciaapi.RestartJobRequest{
		JobId: kusciaAPIJS.jobID,
	})
	assert.Equal(t, res.Data.JobId, kusciaAPIJS.jobID)
}

func TestStopJob(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.AuthRole, consts.AuthRoleMaster)
	ctx = context.WithValue(ctx, consts.SourceDomainKey, "alice")
	res := kusciaAPIJS.StopJob(ctx, &kusciaapi.StopJobRequest{
		JobId: kusciaAPIJS.jobID,
	})
	assert.Equal(t, res.Data.JobId, kusciaAPIJS.jobID)
}

func TestCancelJob(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.AuthRole, consts.AuthRoleMaster)
	ctx = context.WithValue(ctx, consts.SourceDomainKey, "alice")
	res := kusciaAPIJS.CancelJob(ctx, &kusciaapi.CancelJobRequest{
		JobId: kusciaAPIJS.jobID,
	})
	assert.Equal(t, res.Data.JobId, kusciaAPIJS.jobID)
}

func TestBatchQueryJob(t *testing.T) {
	batchResponse := kusciaAPIJS.BatchQueryJobStatus(context.Background(), &kusciaapi.BatchQueryJobStatusRequest{
		JobIds: []string{kusciaAPIJS.jobID},
	})
	assert.Equal(t, len(batchResponse.Data.Jobs), 1)
}

func TestDeleteJob(t *testing.T) {
	deleteRes := kusciaAPIJS.DeleteJob(context.Background(), &kusciaapi.DeleteJobRequest{
		JobId: kusciaAPIJS.jobID,
	})
	assert.Equal(t, deleteRes.Data.JobId, kusciaAPIJS.jobID)
	queryRes := kusciaAPIJS.QueryJob(context.Background(), &kusciaapi.QueryJobRequest{
		JobId: kusciaAPIJS.jobID,
	})
	assert.Equal(t, queryRes.Status.Code, int32(errorcode.ErrorCode_KusciaAPIErrQueryJob))
}
