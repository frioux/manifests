package manifests

func init() {
	paths = map[string]map[string][][]string{
		"io.k8s.api.core.v1.Container": map[string][][]string{
			"io.k8s.api.apps.v1.DaemonSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.DaemonSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.Deployment": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.DeploymentSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.ReplicaSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.StatefulSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.StatefulSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.Deployment": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.StatefulSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DaemonSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.Deployment": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DeploymentSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.StatefulSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1.Job": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1.JobList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1.JobSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1beta1.CronJob": [][]string{
				[]string{"spec", "jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"spec", "jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1beta1.CronJobList": [][]string{
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1beta1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1beta1.JobTemplateSpec": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v2alpha1.CronJob": [][]string{
				[]string{"spec", "jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"spec", "jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v2alpha1.CronJobList": [][]string{
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v2alpha1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v2alpha1.JobTemplateSpec": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.Pod": [][]string{
				[]string{"spec", "containers", ""},
				[]string{"spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.PodList": [][]string{
				[]string{"items", "", "spec", "containers", ""},
				[]string{"items", "", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.PodSpec": [][]string{
				[]string{"containers", ""},
				[]string{"initContainers", ""},
			},
			"io.k8s.api.core.v1.PodTemplate": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.PodTemplateList": [][]string{
				[]string{"items", "", "template", "spec", "containers", ""},
				[]string{"items", "", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.PodTemplateSpec": [][]string{
				[]string{"spec", "containers", ""},
				[]string{"spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.ReplicationController": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.ReplicationControllerList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.ReplicationControllerSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.Deployment": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
		},
		"io.k8s.api.core.v1.PodSpec": map[string][][]string{
			"io.k8s.api.apps.v1.DaemonSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.DaemonSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1.Deployment": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.DeploymentSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.ReplicaSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1.StatefulSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.StatefulSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.Deployment": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.Deployment": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DeploymentSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.batch.v1.Job": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1.JobList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1.JobSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.batch.v1beta1.CronJob": [][]string{
				[]string{"spec", "jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1beta1.CronJobList": [][]string{
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1beta1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1beta1.JobTemplateSpec": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.batch.v2alpha1.CronJob": [][]string{
				[]string{"spec", "jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v2alpha1.CronJobList": [][]string{
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v2alpha1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v2alpha1.JobTemplateSpec": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.core.v1.Pod": [][]string{
				[]string{"spec"},
			},
			"io.k8s.api.core.v1.PodList": [][]string{
				[]string{"items", "", "spec"},
			},
			"io.k8s.api.core.v1.PodTemplate": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.core.v1.PodTemplateList": [][]string{
				[]string{"items", "", "template", "spec"},
			},
			"io.k8s.api.core.v1.PodTemplateSpec": [][]string{
				[]string{"spec"},
			},
			"io.k8s.api.core.v1.ReplicationController": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.core.v1.ReplicationControllerList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.core.v1.ReplicationControllerSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.Deployment": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetSpec": [][]string{
				[]string{"template", "spec"},
			},
		},
		"io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta": map[string][][]string{
			"io.k8s.api.admissionregistration.v1alpha1.InitializerConfiguration": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.admissionregistration.v1alpha1.InitializerConfigurationList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.admissionregistration.v1beta1.MutatingWebhookConfiguration": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.admissionregistration.v1beta1.MutatingWebhookConfigurationList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.admissionregistration.v1beta1.ValidatingWebhookConfiguration": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.admissionregistration.v1beta1.ValidatingWebhookConfigurationList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1.ControllerRevision": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1.ControllerRevisionList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1.DaemonSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.DaemonSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.DaemonSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1.Deployment": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.DeploymentList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.DeploymentSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1.ReplicaSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.ReplicaSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.ReplicaSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1.StatefulSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.StatefulSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.StatefulSetSpec": [][]string{
				[]string{"volumeClaimTemplates", "", "metadata"},
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1beta1.ControllerRevision": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta1.ControllerRevisionList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1beta1.Deployment": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1beta1.Scale": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetSpec": [][]string{
				[]string{"volumeClaimTemplates", "", "metadata"},
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.ControllerRevision": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta2.ControllerRevisionList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.Deployment": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.DeploymentList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.DeploymentSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.Scale": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetSpec": [][]string{
				[]string{"volumeClaimTemplates", "", "metadata"},
				[]string{"template", "metadata"},
			},
			"io.k8s.api.authentication.v1.TokenReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authentication.v1beta1.TokenReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1.LocalSubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1.SelfSubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1.SelfSubjectRulesReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1.SubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1beta1.LocalSubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1beta1.SelfSubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1beta1.SelfSubjectRulesReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1beta1.SubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.autoscaling.v1.HorizontalPodAutoscaler": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.autoscaling.v1.HorizontalPodAutoscalerList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.autoscaling.v1.Scale": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscaler": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscalerList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.batch.v1.Job": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v1.JobList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v1.JobSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.batch.v1beta1.CronJob": [][]string{
				[]string{"metadata"},
				[]string{"spec", "jobTemplate", "metadata"},
				[]string{"spec", "jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v1beta1.CronJobList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "jobTemplate", "metadata"},
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v1beta1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "metadata"},
				[]string{"jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v1beta1.JobTemplateSpec": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v2alpha1.CronJob": [][]string{
				[]string{"metadata"},
				[]string{"spec", "jobTemplate", "metadata"},
				[]string{"spec", "jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v2alpha1.CronJobList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "jobTemplate", "metadata"},
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v2alpha1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "metadata"},
				[]string{"jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v2alpha1.JobTemplateSpec": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.certificates.v1beta1.CertificateSigningRequest": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.certificates.v1beta1.CertificateSigningRequestList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Binding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.ComponentStatus": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.ComponentStatusList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.ConfigMap": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.ConfigMapList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Endpoints": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.EndpointsList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Event": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.EventList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.LimitRange": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.LimitRangeList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Namespace": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.NamespaceList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Node": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.NodeList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.PersistentVolume": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.PersistentVolumeClaim": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.PersistentVolumeClaimList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.PersistentVolumeList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Pod": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.PodList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.PodTemplate": [][]string{
				[]string{"metadata"},
				[]string{"template", "metadata"},
			},
			"io.k8s.api.core.v1.PodTemplateList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "template", "metadata"},
			},
			"io.k8s.api.core.v1.PodTemplateSpec": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.ReplicationController": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.ReplicationControllerList": [][]string{
				[]string{"items", "", "spec", "template", "metadata"},
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.ReplicationControllerSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.core.v1.ResourceQuota": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.ResourceQuotaList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Secret": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.SecretList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Service": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.ServiceAccount": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.ServiceAccountList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.ServiceList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.events.v1beta1.Event": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.events.v1beta1.EventList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "metadata"},
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.Deployment": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "metadata"},
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.Ingress": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.IngressList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.NetworkPolicy": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.NetworkPolicyList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.PodSecurityPolicy": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.PodSecurityPolicyList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "metadata"},
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.Scale": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.networking.v1.NetworkPolicy": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.networking.v1.NetworkPolicyList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.policy.v1beta1.Eviction": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.policy.v1beta1.PodDisruptionBudget": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.policy.v1beta1.PodDisruptionBudgetList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.policy.v1beta1.PodSecurityPolicy": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.policy.v1beta1.PodSecurityPolicyList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1.ClusterRole": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1.ClusterRoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1.ClusterRoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1.ClusterRoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1.Role": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1.RoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1.RoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1.RoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.ClusterRole": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.ClusterRoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.ClusterRoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.ClusterRoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.Role": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.RoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.RoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.RoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1beta1.ClusterRole": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1beta1.ClusterRoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1beta1.ClusterRoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1beta1.ClusterRoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1beta1.Role": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1beta1.RoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1beta1.RoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1beta1.RoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.scheduling.v1alpha1.PriorityClass": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.scheduling.v1alpha1.PriorityClassList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.settings.v1alpha1.PodPreset": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.settings.v1alpha1.PodPresetList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.storage.v1.StorageClass": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.storage.v1.StorageClassList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.storage.v1alpha1.VolumeAttachment": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.storage.v1alpha1.VolumeAttachmentList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.storage.v1beta1.StorageClass": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.storage.v1beta1.StorageClassList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.storage.v1beta1.VolumeAttachment": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.storage.v1beta1.VolumeAttachmentList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1beta1.CustomResourceDefinition": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1beta1.CustomResourceDefinitionList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.APIService": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.APIServiceList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.kube-aggregator.pkg.apis.apiregistration.v1beta1.APIService": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.kube-aggregator.pkg.apis.apiregistration.v1beta1.APIServiceList": [][]string{
				[]string{"items", "", "metadata"},
			},
		},
	}
}

func init() {
	types = map[ktype]string{
		ktype{group: "admissionregistration.k8s.io", kind: "InitializerConfiguration", version: "v1alpha1"}: "io.k8s.api.admissionregistration.v1alpha1.InitializerConfiguration",

		ktype{group: "admissionregistration.k8s.io", kind: "InitializerConfigurationList", version: "v1alpha1"}: "io.k8s.api.admissionregistration.v1alpha1.InitializerConfigurationList",

		ktype{group: "admissionregistration.k8s.io", kind: "MutatingWebhookConfiguration", version: "v1beta1"}: "io.k8s.api.admissionregistration.v1beta1.MutatingWebhookConfiguration",

		ktype{group: "admissionregistration.k8s.io", kind: "MutatingWebhookConfigurationList", version: "v1beta1"}: "io.k8s.api.admissionregistration.v1beta1.MutatingWebhookConfigurationList",

		ktype{group: "admissionregistration.k8s.io", kind: "ValidatingWebhookConfiguration", version: "v1beta1"}: "io.k8s.api.admissionregistration.v1beta1.ValidatingWebhookConfiguration",

		ktype{group: "admissionregistration.k8s.io", kind: "ValidatingWebhookConfigurationList", version: "v1beta1"}: "io.k8s.api.admissionregistration.v1beta1.ValidatingWebhookConfigurationList",

		ktype{group: "apps", kind: "ControllerRevision", version: "v1"}: "io.k8s.api.apps.v1.ControllerRevision",

		ktype{group: "apps", kind: "ControllerRevisionList", version: "v1"}: "io.k8s.api.apps.v1.ControllerRevisionList",

		ktype{group: "apps", kind: "DaemonSet", version: "v1"}: "io.k8s.api.apps.v1.DaemonSet",

		ktype{group: "apps", kind: "DaemonSetList", version: "v1"}: "io.k8s.api.apps.v1.DaemonSetList",

		ktype{group: "apps", kind: "Deployment", version: "v1"}: "io.k8s.api.apps.v1.Deployment",

		ktype{group: "apps", kind: "DeploymentList", version: "v1"}: "io.k8s.api.apps.v1.DeploymentList",

		ktype{group: "apps", kind: "ReplicaSet", version: "v1"}: "io.k8s.api.apps.v1.ReplicaSet",

		ktype{group: "apps", kind: "ReplicaSetList", version: "v1"}: "io.k8s.api.apps.v1.ReplicaSetList",

		ktype{group: "apps", kind: "StatefulSet", version: "v1"}: "io.k8s.api.apps.v1.StatefulSet",

		ktype{group: "apps", kind: "StatefulSetList", version: "v1"}: "io.k8s.api.apps.v1.StatefulSetList",

		ktype{group: "apps", kind: "ControllerRevision", version: "v1beta1"}: "io.k8s.api.apps.v1beta1.ControllerRevision",

		ktype{group: "apps", kind: "ControllerRevisionList", version: "v1beta1"}: "io.k8s.api.apps.v1beta1.ControllerRevisionList",

		ktype{group: "apps", kind: "Deployment", version: "v1beta1"}: "io.k8s.api.apps.v1beta1.Deployment",

		ktype{group: "apps", kind: "DeploymentList", version: "v1beta1"}: "io.k8s.api.apps.v1beta1.DeploymentList",

		ktype{group: "apps", kind: "DeploymentRollback", version: "v1beta1"}: "io.k8s.api.apps.v1beta1.DeploymentRollback",

		ktype{group: "apps", kind: "Scale", version: "v1beta1"}: "io.k8s.api.apps.v1beta1.Scale",

		ktype{group: "apps", kind: "StatefulSet", version: "v1beta1"}: "io.k8s.api.apps.v1beta1.StatefulSet",

		ktype{group: "apps", kind: "StatefulSetList", version: "v1beta1"}: "io.k8s.api.apps.v1beta1.StatefulSetList",

		ktype{group: "apps", kind: "ControllerRevision", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.ControllerRevision",

		ktype{group: "apps", kind: "ControllerRevisionList", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.ControllerRevisionList",

		ktype{group: "apps", kind: "DaemonSet", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.DaemonSet",

		ktype{group: "apps", kind: "DaemonSetList", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.DaemonSetList",

		ktype{group: "apps", kind: "Deployment", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.Deployment",

		ktype{group: "apps", kind: "DeploymentList", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.DeploymentList",

		ktype{group: "apps", kind: "ReplicaSet", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.ReplicaSet",

		ktype{group: "apps", kind: "ReplicaSetList", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.ReplicaSetList",

		ktype{group: "apps", kind: "Scale", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.Scale",

		ktype{group: "apps", kind: "StatefulSet", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.StatefulSet",

		ktype{group: "apps", kind: "StatefulSetList", version: "v1beta2"}: "io.k8s.api.apps.v1beta2.StatefulSetList",

		ktype{group: "authentication.k8s.io", kind: "TokenReview", version: "v1"}: "io.k8s.api.authentication.v1.TokenReview",

		ktype{group: "authentication.k8s.io", kind: "TokenReview", version: "v1beta1"}: "io.k8s.api.authentication.v1beta1.TokenReview",

		ktype{group: "authorization.k8s.io", kind: "LocalSubjectAccessReview", version: "v1"}: "io.k8s.api.authorization.v1.LocalSubjectAccessReview",

		ktype{group: "authorization.k8s.io", kind: "SelfSubjectAccessReview", version: "v1"}: "io.k8s.api.authorization.v1.SelfSubjectAccessReview",

		ktype{group: "authorization.k8s.io", kind: "SelfSubjectRulesReview", version: "v1"}: "io.k8s.api.authorization.v1.SelfSubjectRulesReview",

		ktype{group: "authorization.k8s.io", kind: "SubjectAccessReview", version: "v1"}: "io.k8s.api.authorization.v1.SubjectAccessReview",

		ktype{group: "authorization.k8s.io", kind: "LocalSubjectAccessReview", version: "v1beta1"}: "io.k8s.api.authorization.v1beta1.LocalSubjectAccessReview",

		ktype{group: "authorization.k8s.io", kind: "SelfSubjectAccessReview", version: "v1beta1"}: "io.k8s.api.authorization.v1beta1.SelfSubjectAccessReview",

		ktype{group: "authorization.k8s.io", kind: "SelfSubjectRulesReview", version: "v1beta1"}: "io.k8s.api.authorization.v1beta1.SelfSubjectRulesReview",

		ktype{group: "authorization.k8s.io", kind: "SubjectAccessReview", version: "v1beta1"}: "io.k8s.api.authorization.v1beta1.SubjectAccessReview",

		ktype{group: "autoscaling", kind: "HorizontalPodAutoscaler", version: "v1"}: "io.k8s.api.autoscaling.v1.HorizontalPodAutoscaler",

		ktype{group: "autoscaling", kind: "HorizontalPodAutoscalerList", version: "v1"}: "io.k8s.api.autoscaling.v1.HorizontalPodAutoscalerList",

		ktype{group: "autoscaling", kind: "Scale", version: "v1"}: "io.k8s.api.autoscaling.v1.Scale",

		ktype{group: "autoscaling", kind: "HorizontalPodAutoscaler", version: "v2beta1"}: "io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscaler",

		ktype{group: "autoscaling", kind: "HorizontalPodAutoscalerList", version: "v2beta1"}: "io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscalerList",

		ktype{group: "batch", kind: "Job", version: "v1"}: "io.k8s.api.batch.v1.Job",

		ktype{group: "batch", kind: "JobList", version: "v1"}: "io.k8s.api.batch.v1.JobList",

		ktype{group: "batch", kind: "CronJob", version: "v1beta1"}: "io.k8s.api.batch.v1beta1.CronJob",

		ktype{group: "batch", kind: "CronJobList", version: "v1beta1"}: "io.k8s.api.batch.v1beta1.CronJobList",

		ktype{group: "batch", kind: "CronJob", version: "v2alpha1"}: "io.k8s.api.batch.v2alpha1.CronJob",

		ktype{group: "batch", kind: "CronJobList", version: "v2alpha1"}: "io.k8s.api.batch.v2alpha1.CronJobList",

		ktype{group: "certificates.k8s.io", kind: "CertificateSigningRequest", version: "v1beta1"}: "io.k8s.api.certificates.v1beta1.CertificateSigningRequest",

		ktype{group: "certificates.k8s.io", kind: "CertificateSigningRequestList", version: "v1beta1"}: "io.k8s.api.certificates.v1beta1.CertificateSigningRequestList",

		ktype{group: "", kind: "Binding", version: "v1"}: "io.k8s.api.core.v1.Binding",

		ktype{group: "", kind: "ComponentStatus", version: "v1"}: "io.k8s.api.core.v1.ComponentStatus",

		ktype{group: "", kind: "ComponentStatusList", version: "v1"}: "io.k8s.api.core.v1.ComponentStatusList",

		ktype{group: "", kind: "ConfigMap", version: "v1"}: "io.k8s.api.core.v1.ConfigMap",

		ktype{group: "", kind: "ConfigMapList", version: "v1"}: "io.k8s.api.core.v1.ConfigMapList",

		ktype{group: "", kind: "Endpoints", version: "v1"}: "io.k8s.api.core.v1.Endpoints",

		ktype{group: "", kind: "EndpointsList", version: "v1"}: "io.k8s.api.core.v1.EndpointsList",

		ktype{group: "", kind: "Event", version: "v1"}: "io.k8s.api.core.v1.Event",

		ktype{group: "", kind: "EventList", version: "v1"}: "io.k8s.api.core.v1.EventList",

		ktype{group: "", kind: "LimitRange", version: "v1"}: "io.k8s.api.core.v1.LimitRange",

		ktype{group: "", kind: "LimitRangeList", version: "v1"}: "io.k8s.api.core.v1.LimitRangeList",

		ktype{group: "", kind: "Namespace", version: "v1"}: "io.k8s.api.core.v1.Namespace",

		ktype{group: "", kind: "NamespaceList", version: "v1"}: "io.k8s.api.core.v1.NamespaceList",

		ktype{group: "", kind: "Node", version: "v1"}: "io.k8s.api.core.v1.Node",

		ktype{group: "", kind: "NodeConfigSource", version: "v1"}: "io.k8s.api.core.v1.NodeConfigSource",

		ktype{group: "", kind: "NodeList", version: "v1"}: "io.k8s.api.core.v1.NodeList",

		ktype{group: "", kind: "PersistentVolume", version: "v1"}: "io.k8s.api.core.v1.PersistentVolume",

		ktype{group: "", kind: "PersistentVolumeClaim", version: "v1"}: "io.k8s.api.core.v1.PersistentVolumeClaim",

		ktype{group: "", kind: "PersistentVolumeClaimList", version: "v1"}: "io.k8s.api.core.v1.PersistentVolumeClaimList",

		ktype{group: "", kind: "PersistentVolumeList", version: "v1"}: "io.k8s.api.core.v1.PersistentVolumeList",

		ktype{group: "", kind: "Pod", version: "v1"}: "io.k8s.api.core.v1.Pod",

		ktype{group: "", kind: "PodList", version: "v1"}: "io.k8s.api.core.v1.PodList",

		ktype{group: "", kind: "PodTemplate", version: "v1"}: "io.k8s.api.core.v1.PodTemplate",

		ktype{group: "", kind: "PodTemplateList", version: "v1"}: "io.k8s.api.core.v1.PodTemplateList",

		ktype{group: "", kind: "ReplicationController", version: "v1"}: "io.k8s.api.core.v1.ReplicationController",

		ktype{group: "", kind: "ReplicationControllerList", version: "v1"}: "io.k8s.api.core.v1.ReplicationControllerList",

		ktype{group: "", kind: "ResourceQuota", version: "v1"}: "io.k8s.api.core.v1.ResourceQuota",

		ktype{group: "", kind: "ResourceQuotaList", version: "v1"}: "io.k8s.api.core.v1.ResourceQuotaList",

		ktype{group: "", kind: "Secret", version: "v1"}: "io.k8s.api.core.v1.Secret",

		ktype{group: "", kind: "SecretList", version: "v1"}: "io.k8s.api.core.v1.SecretList",

		ktype{group: "", kind: "Service", version: "v1"}: "io.k8s.api.core.v1.Service",

		ktype{group: "", kind: "ServiceAccount", version: "v1"}: "io.k8s.api.core.v1.ServiceAccount",

		ktype{group: "", kind: "ServiceAccountList", version: "v1"}: "io.k8s.api.core.v1.ServiceAccountList",

		ktype{group: "", kind: "ServiceList", version: "v1"}: "io.k8s.api.core.v1.ServiceList",

		ktype{group: "events.k8s.io", kind: "Event", version: "v1beta1"}: "io.k8s.api.events.v1beta1.Event",

		ktype{group: "events.k8s.io", kind: "EventList", version: "v1beta1"}: "io.k8s.api.events.v1beta1.EventList",

		ktype{group: "extensions", kind: "DaemonSet", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.DaemonSet",

		ktype{group: "extensions", kind: "DaemonSetList", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.DaemonSetList",

		ktype{group: "extensions", kind: "Deployment", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.Deployment",

		ktype{group: "extensions", kind: "DeploymentList", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.DeploymentList",

		ktype{group: "extensions", kind: "DeploymentRollback", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.DeploymentRollback",

		ktype{group: "extensions", kind: "Ingress", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.Ingress",

		ktype{group: "extensions", kind: "IngressList", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.IngressList",

		ktype{group: "extensions", kind: "NetworkPolicy", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.NetworkPolicy",

		ktype{group: "extensions", kind: "NetworkPolicyList", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.NetworkPolicyList",

		ktype{group: "extensions", kind: "PodSecurityPolicy", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.PodSecurityPolicy",

		ktype{group: "extensions", kind: "PodSecurityPolicyList", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.PodSecurityPolicyList",

		ktype{group: "extensions", kind: "ReplicaSet", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.ReplicaSet",

		ktype{group: "extensions", kind: "ReplicaSetList", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.ReplicaSetList",

		ktype{group: "extensions", kind: "Scale", version: "v1beta1"}: "io.k8s.api.extensions.v1beta1.Scale",

		ktype{group: "networking.k8s.io", kind: "NetworkPolicy", version: "v1"}: "io.k8s.api.networking.v1.NetworkPolicy",

		ktype{group: "networking.k8s.io", kind: "NetworkPolicyList", version: "v1"}: "io.k8s.api.networking.v1.NetworkPolicyList",

		ktype{group: "policy", kind: "Eviction", version: "v1beta1"}: "io.k8s.api.policy.v1beta1.Eviction",

		ktype{group: "policy", kind: "PodDisruptionBudget", version: "v1beta1"}: "io.k8s.api.policy.v1beta1.PodDisruptionBudget",

		ktype{group: "policy", kind: "PodDisruptionBudgetList", version: "v1beta1"}: "io.k8s.api.policy.v1beta1.PodDisruptionBudgetList",

		ktype{group: "policy", kind: "PodSecurityPolicy", version: "v1beta1"}: "io.k8s.api.policy.v1beta1.PodSecurityPolicy",

		ktype{group: "policy", kind: "PodSecurityPolicyList", version: "v1beta1"}: "io.k8s.api.policy.v1beta1.PodSecurityPolicyList",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRole", version: "v1"}: "io.k8s.api.rbac.v1.ClusterRole",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRoleBinding", version: "v1"}: "io.k8s.api.rbac.v1.ClusterRoleBinding",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRoleBindingList", version: "v1"}: "io.k8s.api.rbac.v1.ClusterRoleBindingList",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRoleList", version: "v1"}: "io.k8s.api.rbac.v1.ClusterRoleList",

		ktype{group: "rbac.authorization.k8s.io", kind: "Role", version: "v1"}: "io.k8s.api.rbac.v1.Role",

		ktype{group: "rbac.authorization.k8s.io", kind: "RoleBinding", version: "v1"}: "io.k8s.api.rbac.v1.RoleBinding",

		ktype{group: "rbac.authorization.k8s.io", kind: "RoleBindingList", version: "v1"}: "io.k8s.api.rbac.v1.RoleBindingList",

		ktype{group: "rbac.authorization.k8s.io", kind: "RoleList", version: "v1"}: "io.k8s.api.rbac.v1.RoleList",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRole", version: "v1alpha1"}: "io.k8s.api.rbac.v1alpha1.ClusterRole",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRoleBinding", version: "v1alpha1"}: "io.k8s.api.rbac.v1alpha1.ClusterRoleBinding",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRoleBindingList", version: "v1alpha1"}: "io.k8s.api.rbac.v1alpha1.ClusterRoleBindingList",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRoleList", version: "v1alpha1"}: "io.k8s.api.rbac.v1alpha1.ClusterRoleList",

		ktype{group: "rbac.authorization.k8s.io", kind: "Role", version: "v1alpha1"}: "io.k8s.api.rbac.v1alpha1.Role",

		ktype{group: "rbac.authorization.k8s.io", kind: "RoleBinding", version: "v1alpha1"}: "io.k8s.api.rbac.v1alpha1.RoleBinding",

		ktype{group: "rbac.authorization.k8s.io", kind: "RoleBindingList", version: "v1alpha1"}: "io.k8s.api.rbac.v1alpha1.RoleBindingList",

		ktype{group: "rbac.authorization.k8s.io", kind: "RoleList", version: "v1alpha1"}: "io.k8s.api.rbac.v1alpha1.RoleList",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRole", version: "v1beta1"}: "io.k8s.api.rbac.v1beta1.ClusterRole",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRoleBinding", version: "v1beta1"}: "io.k8s.api.rbac.v1beta1.ClusterRoleBinding",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRoleBindingList", version: "v1beta1"}: "io.k8s.api.rbac.v1beta1.ClusterRoleBindingList",

		ktype{group: "rbac.authorization.k8s.io", kind: "ClusterRoleList", version: "v1beta1"}: "io.k8s.api.rbac.v1beta1.ClusterRoleList",

		ktype{group: "rbac.authorization.k8s.io", kind: "Role", version: "v1beta1"}: "io.k8s.api.rbac.v1beta1.Role",

		ktype{group: "rbac.authorization.k8s.io", kind: "RoleBinding", version: "v1beta1"}: "io.k8s.api.rbac.v1beta1.RoleBinding",

		ktype{group: "rbac.authorization.k8s.io", kind: "RoleBindingList", version: "v1beta1"}: "io.k8s.api.rbac.v1beta1.RoleBindingList",

		ktype{group: "rbac.authorization.k8s.io", kind: "RoleList", version: "v1beta1"}: "io.k8s.api.rbac.v1beta1.RoleList",

		ktype{group: "scheduling.k8s.io", kind: "PriorityClass", version: "v1alpha1"}: "io.k8s.api.scheduling.v1alpha1.PriorityClass",

		ktype{group: "scheduling.k8s.io", kind: "PriorityClassList", version: "v1alpha1"}: "io.k8s.api.scheduling.v1alpha1.PriorityClassList",

		ktype{group: "settings.k8s.io", kind: "PodPreset", version: "v1alpha1"}: "io.k8s.api.settings.v1alpha1.PodPreset",

		ktype{group: "settings.k8s.io", kind: "PodPresetList", version: "v1alpha1"}: "io.k8s.api.settings.v1alpha1.PodPresetList",

		ktype{group: "storage.k8s.io", kind: "StorageClass", version: "v1"}: "io.k8s.api.storage.v1.StorageClass",

		ktype{group: "storage.k8s.io", kind: "StorageClassList", version: "v1"}: "io.k8s.api.storage.v1.StorageClassList",

		ktype{group: "storage.k8s.io", kind: "VolumeAttachment", version: "v1alpha1"}: "io.k8s.api.storage.v1alpha1.VolumeAttachment",

		ktype{group: "storage.k8s.io", kind: "VolumeAttachmentList", version: "v1alpha1"}: "io.k8s.api.storage.v1alpha1.VolumeAttachmentList",

		ktype{group: "storage.k8s.io", kind: "StorageClass", version: "v1beta1"}: "io.k8s.api.storage.v1beta1.StorageClass",

		ktype{group: "storage.k8s.io", kind: "StorageClassList", version: "v1beta1"}: "io.k8s.api.storage.v1beta1.StorageClassList",

		ktype{group: "storage.k8s.io", kind: "VolumeAttachment", version: "v1beta1"}: "io.k8s.api.storage.v1beta1.VolumeAttachment",

		ktype{group: "storage.k8s.io", kind: "VolumeAttachmentList", version: "v1beta1"}: "io.k8s.api.storage.v1beta1.VolumeAttachmentList",

		ktype{group: "", kind: "APIGroup", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.APIGroup",

		ktype{group: "", kind: "APIGroupList", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.APIGroupList",

		ktype{group: "", kind: "APIResourceList", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.APIResourceList",

		ktype{group: "", kind: "APIVersions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.APIVersions",

		ktype{group: "", kind: "DeleteOptions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "admission.k8s.io", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "admissionregistration.k8s.io", kind: "DeleteOptions", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "admissionregistration.k8s.io", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "apps", kind: "DeleteOptions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "apps", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "apps", kind: "DeleteOptions", version: "v1beta2"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "authentication.k8s.io", kind: "DeleteOptions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "authentication.k8s.io", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "authorization.k8s.io", kind: "DeleteOptions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "authorization.k8s.io", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "autoscaling", kind: "DeleteOptions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "autoscaling", kind: "DeleteOptions", version: "v2beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "batch", kind: "DeleteOptions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "batch", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "batch", kind: "DeleteOptions", version: "v2alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "certificates.k8s.io", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "events.k8s.io", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "extensions", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "imagepolicy.k8s.io", kind: "DeleteOptions", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "networking.k8s.io", kind: "DeleteOptions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "policy", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "rbac.authorization.k8s.io", kind: "DeleteOptions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "rbac.authorization.k8s.io", kind: "DeleteOptions", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "rbac.authorization.k8s.io", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "scheduling.k8s.io", kind: "DeleteOptions", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "settings.k8s.io", kind: "DeleteOptions", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "storage.k8s.io", kind: "DeleteOptions", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "storage.k8s.io", kind: "DeleteOptions", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",
		ktype{group: "storage.k8s.io", kind: "DeleteOptions", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions",

		ktype{group: "", kind: "Status", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.Status",

		ktype{group: "", kind: "WatchEvent", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "admission.k8s.io", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "admissionregistration.k8s.io", kind: "WatchEvent", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "admissionregistration.k8s.io", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "apps", kind: "WatchEvent", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "apps", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "apps", kind: "WatchEvent", version: "v1beta2"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "authentication.k8s.io", kind: "WatchEvent", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "authentication.k8s.io", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "authorization.k8s.io", kind: "WatchEvent", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "authorization.k8s.io", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "autoscaling", kind: "WatchEvent", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "autoscaling", kind: "WatchEvent", version: "v2beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "batch", kind: "WatchEvent", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "batch", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "batch", kind: "WatchEvent", version: "v2alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "certificates.k8s.io", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "events.k8s.io", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "extensions", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "imagepolicy.k8s.io", kind: "WatchEvent", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "networking.k8s.io", kind: "WatchEvent", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "policy", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "rbac.authorization.k8s.io", kind: "WatchEvent", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "rbac.authorization.k8s.io", kind: "WatchEvent", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "rbac.authorization.k8s.io", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "scheduling.k8s.io", kind: "WatchEvent", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "settings.k8s.io", kind: "WatchEvent", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "storage.k8s.io", kind: "WatchEvent", version: "v1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "storage.k8s.io", kind: "WatchEvent", version: "v1alpha1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",
		ktype{group: "storage.k8s.io", kind: "WatchEvent", version: "v1beta1"}: "io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent",

	}
}
