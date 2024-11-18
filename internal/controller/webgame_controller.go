/*
Copyright 2024.

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

package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	webgamev1 "github.com/ambition0222/webgame/api/v1"
)

// WebGameReconciler reconciles a WebGame object
type WebGameReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=webgame.webgame.tech,resources=webgames,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=webgame.webgame.tech,resources=webgames/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=webgame.webgame.tech,resources=webgames/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the WebGame object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile
func (r *WebGameReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	// 接收事件
	l.Info("webgame event received")
	// 完成事件处理
	defer func() { l.Info("webgame event handling completed") }()

	// 获取webgame实例
	var webgame webgamev1.WebGame
	if err := r.Get(ctx, req.NamespacedName, &webgame); err != nil {
		// 不存在,触发的是删除事件
		if errors.IsNotFound(err) {
			l.Info("webgame not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// 打印日志，等待重试
		l.Error(err, "unable to fetch webgame,requeue")
	}
	// 成功获取，打印日志
	l.Info("show webgame name", "name", webgame.GetName())
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WebGameReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webgamev1.WebGame{}).
		Named("webgame").
		Complete(r)
}
