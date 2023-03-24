package analyzer

import (
	"context"

	"github.com/k8sgpt-ai/k8sgpt/pkg/ai"
	"github.com/k8sgpt-ai/k8sgpt/pkg/kubernetes"
)

func RunAnalysis(ctx context.Context, client *kubernetes.Client, aiClient ai.IAI, explain bool) error {
	err := AnalyzePod(ctx, client, aiClient, explain)
	if err != nil {
		return err
	}

	err = AnalyzeReplicaSet(ctx, client, aiClient, explain)
	if err != nil {
		return err
	}
	return nil
}
