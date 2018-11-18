package manifests

func init() {
	paths = map[string]map[string][][]string{
		"io.k8s.api.core.v1.Container": map[string][][]string{
			"io.k8s.api.extensions.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1beta1.CronJobList": [][]string{
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.DaemonSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.Deployment": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.DaemonSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.PodList": [][]string{
				[]string{"items", "", "spec", "containers", ""},
				[]string{"items", "", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.PodTemplateSpec": [][]string{
				[]string{"spec", "containers", ""},
				[]string{"spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.Pod": [][]string{
				[]string{"spec", "containers", ""},
				[]string{"spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v2alpha1.CronJob": [][]string{
				[]string{"spec", "jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"spec", "jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.PodSpec": [][]string{
				[]string{"containers", ""},
				[]string{"initContainers", ""},
			},
			"io.k8s.api.batch.v2alpha1.JobTemplateSpec": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DeploymentSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v2alpha1.CronJobList": [][]string{
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.StatefulSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1beta1.JobTemplateSpec": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.PodTemplateList": [][]string{
				[]string{"items", "", "template", "spec", "containers", ""},
				[]string{"items", "", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.Deployment": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1beta1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.StatefulSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1beta1.CronJob": [][]string{
				[]string{"spec", "jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"spec", "jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1.JobSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.ReplicationControllerSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.DeploymentSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1.JobList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.StatefulSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.ReplicationController": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.PodTemplate": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.ReplicaSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.Deployment": [][]string{
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
			"io.k8s.api.apps.v1.StatefulSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.core.v1.ReplicationControllerList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.Deployment": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v1.Job": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.DaemonSet": [][]string{
				[]string{"spec", "template", "spec", "containers", ""},
				[]string{"spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.batch.v2alpha1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "spec", "template", "spec", "containers", ""},
				[]string{"jobTemplate", "spec", "template", "spec", "initContainers", ""},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetSpec": [][]string{
				[]string{"template", "spec", "containers", ""},
				[]string{"template", "spec", "initContainers", ""},
			},
			"io.k8s.api.apps.v1.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec", "containers", ""},
				[]string{"items", "", "spec", "template", "spec", "initContainers", ""},
			},
		},
		"io.k8s.api.core.v1.PodSpec": map[string][][]string{
			"io.k8s.api.extensions.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1beta1.CronJobList": [][]string{
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.DaemonSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.Deployment": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1.DaemonSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.core.v1.PodList": [][]string{
				[]string{"items", "", "spec"},
			},
			"io.k8s.api.core.v1.PodTemplateSpec": [][]string{
				[]string{"spec"},
			},
			"io.k8s.api.core.v1.Pod": [][]string{
				[]string{"spec"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v2alpha1.CronJob": [][]string{
				[]string{"spec", "jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v2alpha1.JobTemplateSpec": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.batch.v2alpha1.CronJobList": [][]string{
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DeploymentSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1beta1.JobTemplateSpec": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.Deployment": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.core.v1.PodTemplateList": [][]string{
				[]string{"items", "", "template", "spec"},
			},
			"io.k8s.api.batch.v1beta1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1beta1.CronJob": [][]string{
				[]string{"spec", "jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1.JobSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.core.v1.ReplicationControllerSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.DeploymentSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.batch.v1.JobList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.StatefulSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.core.v1.ReplicationController": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.core.v1.PodTemplate": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1.ReplicaSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.Deployment": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.core.v1.ReplicationControllerList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1.StatefulSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.Deployment": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.batch.v1.Job": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSet": [][]string{
				[]string{"spec", "template", "spec"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
			"io.k8s.api.batch.v2alpha1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "spec", "template", "spec"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetSpec": [][]string{
				[]string{"template", "spec"},
			},
			"io.k8s.api.apps.v1.StatefulSetList": [][]string{
				[]string{"items", "", "spec", "template", "spec"},
			},
		},
		"io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta": map[string][][]string{
			"io.k8s.api.storage.v1beta1.VolumeAttachment": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta1.ControllerRevisionList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.admissionregistration.v1beta1.MutatingWebhookConfigurationList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.batch.v1beta1.CronJob": [][]string{
				[]string{"metadata"},
				[]string{"spec", "jobTemplate", "metadata"},
				[]string{"spec", "jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.EventList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.storage.v1beta1.StorageClass": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.Ingress": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.storage.v1beta1.VolumeAttachmentList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Node": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.certificates.v1beta1.CertificateSigningRequest": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta1.ControllerRevision": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.Endpoints": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.PodSecurityPolicy": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.policy.v1beta1.Eviction": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1beta1.SelfSubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta2.Scale": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.scheduling.v1alpha1.PriorityClass": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.settings.v1alpha1.PodPreset": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.ClusterRoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.authorization.v1beta1.SubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.autoscaling.v1.Scale": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.Event": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.EndpointsList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1beta1.ClusterRoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1.DaemonSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetList": [][]string{
				[]string{"items", "", "spec", "template", "metadata"},
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.networking.v1.NetworkPolicyList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Namespace": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.admissionregistration.v1alpha1.InitializerConfigurationList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Binding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.Secret": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1.DaemonSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.NetworkPolicyList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.PersistentVolume": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.admissionregistration.v1beta1.MutatingWebhookConfiguration": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.PersistentVolumeClaim": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.ResourceQuota": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.ClusterRoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.events.v1beta1.Event": [][]string{
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
			"io.k8s.api.rbac.v1alpha1.ClusterRoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.batch.v1beta1.JobTemplateSpec": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.authentication.v1.TokenReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.policy.v1beta1.PodSecurityPolicyList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1.DeploymentSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1.DeploymentList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.StatefulSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.rbac.v1.RoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.ReplicationController": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.rbac.v1beta1.RoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1.ClusterRoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.ServiceAccount": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.rbac.v1.RoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1.StatefulSetSpec": [][]string{
				[]string{"volumeClaimTemplates", "", "metadata"},
				[]string{"template", "metadata"},
			},
			"io.k8s.api.core.v1.ServiceAccountList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1beta1.ClusterRoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.RoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.ControllerRevisionList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.Deployment": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.authorization.v1.SubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.batch.v2alpha1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "metadata"},
				[]string{"jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1.StatefulSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.events.v1beta1.EventList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.authentication.v1beta1.TokenReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1.Role": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.policy.v1beta1.PodDisruptionBudget": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1.ClusterRoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.admissionregistration.v1beta1.ValidatingWebhookConfiguration": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.PodTemplateSpec": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1beta1.CustomResourceDefinition": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.APIServiceList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.batch.v2alpha1.JobTemplateSpec": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.NodeList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscaler": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.RoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.kube-aggregator.pkg.apis.apiregistration.v1beta1.APIService": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.storage.v1alpha1.VolumeAttachment": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1beta1.ClusterRoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1.RoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.authorization.v1.LocalSubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.authorization.v1beta1.LocalSubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.ReplicationControllerSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.core.v1.ResourceQuotaList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.admissionregistration.v1alpha1.InitializerConfiguration": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DaemonSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.core.v1.ConfigMapList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.policy.v1beta1.PodDisruptionBudgetList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.storage.v1.StorageClass": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1beta1.SelfSubjectRulesReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta2.DeploymentList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.PodTemplate": [][]string{
				[]string{"metadata"},
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1.ReplicaSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.Deployment": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.ReplicationControllerList": [][]string{
				[]string{"items", "", "spec", "template", "metadata"},
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.authorization.v1.SelfSubjectRulesReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.ComponentStatus": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscalerList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.ClusterRole": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.settings.v1alpha1.PodPresetList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.batch.v1.Job": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.certificates.v1beta1.CertificateSigningRequestList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.storage.v1alpha1.VolumeAttachmentList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.PersistentVolumeList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1.ControllerRevisionList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.batch.v1beta1.CronJobList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "jobTemplate", "metadata"},
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1beta1.CustomResourceDefinitionList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1beta1.Deployment": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.DaemonSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.core.v1.SecretList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetList": [][]string{
				[]string{"items", "", "spec", "template", "metadata"},
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.Pod": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.batch.v2alpha1.CronJob": [][]string{
				[]string{"metadata"},
				[]string{"spec", "jobTemplate", "metadata"},
				[]string{"spec", "jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.IngressList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.ServiceList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.autoscaling.v1.HorizontalPodAutoscaler": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1.ReplicaSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.scheduling.v1alpha1.PriorityClassList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.PodSecurityPolicyList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.policy.v1beta1.PodSecurityPolicy": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1.ControllerRevision": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1.Deployment": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.batch.v1beta1.CronJobSpec": [][]string{
				[]string{"jobTemplate", "metadata"},
				[]string{"jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.rbac.v1beta1.Role": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.Service": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.LimitRangeList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.ControllerRevision": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1.DaemonSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.ComponentStatusList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.batch.v1.JobSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.batch.v1.JobList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.rbac.v1.ClusterRoleBinding": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.Scale": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.authorization.v1.SelfSubjectAccessReview": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.autoscaling.v1.HorizontalPodAutoscalerList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.PersistentVolumeClaimList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.admissionregistration.v1beta1.ValidatingWebhookConfigurationList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.Role": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1beta1.ClusterRole": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.apps.v1beta1.Scale": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1beta1.RoleList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1beta1.StatefulSetSpec": [][]string{
				[]string{"volumeClaimTemplates", "", "metadata"},
				[]string{"template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetSpec": [][]string{
				[]string{"volumeClaimTemplates", "", "metadata"},
				[]string{"template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSetSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.DeploymentList": [][]string{
				[]string{"items", "", "spec", "template", "metadata"},
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.storage.v1beta1.StorageClassList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1.ReplicaSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.PodList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1alpha1.RoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.extensions.v1beta1.NetworkPolicy": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.rbac.v1.ClusterRole": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.APIService": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.extensions.v1beta1.ReplicaSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.ReplicaSet": [][]string{
				[]string{"metadata"},
				[]string{"spec", "template", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.DeploymentSpec": [][]string{
				[]string{"template", "metadata"},
			},
			"io.k8s.api.batch.v2alpha1.CronJobList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "jobTemplate", "metadata"},
				[]string{"items", "", "spec", "jobTemplate", "spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.LimitRange": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.networking.v1.NetworkPolicy": [][]string{
				[]string{"metadata"},
			},
			"io.k8s.api.core.v1.PodTemplateList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "template", "metadata"},
			},
			"io.k8s.kube-aggregator.pkg.apis.apiregistration.v1beta1.APIServiceList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.apps.v1beta2.StatefulSetList": [][]string{
				[]string{"items", "", "metadata"},
				[]string{"items", "", "spec", "volumeClaimTemplates", "", "metadata"},
				[]string{"items", "", "spec", "template", "metadata"},
			},
			"io.k8s.api.core.v1.NamespaceList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.storage.v1.StorageClassList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.rbac.v1beta1.RoleBindingList": [][]string{
				[]string{"items", "", "metadata"},
			},
			"io.k8s.api.core.v1.ConfigMap": [][]string{
				[]string{"metadata"},
			},
		},
	}
}
