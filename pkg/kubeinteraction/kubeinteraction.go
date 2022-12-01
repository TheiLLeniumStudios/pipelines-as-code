package kubeinteraction

import (
	"context"

	"github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode/v1alpha1"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/params"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"

	ktypes "github.com/openshift-pipelines/pipelines-as-code/pkg/secrets/types"
)

type Interface interface {
	CleanupPipelines(context.Context, *zap.SugaredLogger, *v1alpha1.Repository, *v1beta1.PipelineRun, int) error
	CreateSecret(ctx context.Context, ns string, secret *corev1.Secret) error
	DeleteSecret(context.Context, *zap.SugaredLogger, string, string) error
	GetSecret(context.Context, ktypes.GetSecretOpt) (string, error)
	GetPodLogs(context.Context, string, string, string, int64) (string, error)
}

type Interaction struct {
	Run *params.Run
}

// validate the interface implementation
var _ Interface = (*Interaction)(nil)

func NewKubernetesInteraction(c *params.Run) (*Interaction, error) {
	return &Interaction{
		Run: c,
	}, nil
}
