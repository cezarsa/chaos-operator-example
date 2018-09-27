package stub

import (
	"context"
	"log"
	"math/rand"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"

	"github.com/cezarsa/chaos/pkg/apis/chaos/v1beta1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1beta1.PodChaos:
		now := time.Now()
		if time.Since(o.Status.LastRun) < time.Duration(o.Spec.FrequencySeconds)*time.Second {
			return nil
		}

		defer func() {
			o.Status.LastRun = now
			sdk.Update(o)
		}()

		podList := corev1.PodList{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1",
			},
		}

		listOpts := sdk.WithListOptions(&metav1.ListOptions{
			LabelSelector: labels.SelectorFromSet(labels.Set(o.Spec.Selector)).String(),
		})

		if err := sdk.List(o.Namespace, &podList, listOpts); err != nil {
			return err
		}

		if len(podList.Items) == 0 {
			return nil
		}

		victim := podList.Items[rand.Intn(len(podList.Items))]

		log.Printf("Chaos victim: %v\n", victim.Name)

		return sdk.Delete(&victim)

	}
	return nil
}
