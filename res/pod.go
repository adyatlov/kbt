package res

import (
	"time"
)
type Pods struct {
	Items []Pod
}
type Pod struct {
APIVersion string   `yaml:"apiVersion"`
Kind     string   `yaml:"kind"`
Metadata Metadata `yaml:"metadata"`
Spec     Spec     `yaml:"spec"`
Status   Status   `yaml:"status"`
}
type OwnerReferences struct {
APIVersion         string `yaml:"apiVersion"`
BlockOwnerDeletion bool   `yaml:"blockOwnerDeletion"`
Controller         bool   `yaml:"controller"`
Kind               string `yaml:"kind"`
Name               string `yaml:"name"`
UID                string `yaml:"uid"`
}
type Metadata struct {
CreationTimestamp time.Time         `yaml:"creationTimestamp"`
GenerateName      string            `yaml:"generateName"`
Name              string            `yaml:"name"`
Namespace         string            `yaml:"namespace"`
OwnerReferences   []OwnerReferences `yaml:"ownerReferences"`
ResourceVersion   string            `yaml:"resourceVersion"`
SelfLink          string            `yaml:"selfLink"`
UID               string            `yaml:"uid"`
}
type Limits struct {
CPU    string `yaml:"cpu"`
Memory string `yaml:"memory"`
}
type Requests struct {
CPU    string `yaml:"cpu"`
Memory string `yaml:"memory"`
}
type Resources struct {
Limits   Limits   `yaml:"limits"`
Requests Requests `yaml:"requests"`
}
type VolumeMounts struct {
MountPath string `yaml:"mountPath"`
Name      string `yaml:"name"`
ReadOnly  bool   `yaml:"readOnly"`
}
type Container struct {
Command                  []string       `yaml:"command"`
Image                    string         `yaml:"image"`
ImagePullPolicy          string         `yaml:"imagePullPolicy"`
Name                     string         `yaml:"name"`
Resources                Resources      `yaml:"resources"`
TerminationMessagePath   string         `yaml:"terminationMessagePath"`
TerminationMessagePolicy string         `yaml:"terminationMessagePolicy"`
VolumeMounts             []VolumeMounts `yaml:"volumeMounts"`
}
type SecurityContext struct {
}
type Toleration struct {
Effect            string `yaml:"effect"`
Key               string `yaml:"key"`
Operator          string `yaml:"operator"`
TolerationSeconds int    `yaml:"tolerationSeconds"`
}
type Secret struct {
DefaultMode int    `yaml:"defaultMode"`
SecretName  string `yaml:"secretName"`
}
type Volumes struct {
Name   string `yaml:"name"`
Secret Secret `yaml:"secret"`
}
type Spec struct {
Containers                    []Container     `yaml:"containers"`
DNSPolicy                     string          `yaml:"dnsPolicy"`
EnableServiceLinks            bool            `yaml:"enableServiceLinks"`
NodeName                      string          `yaml:"nodeName"`
PreemptionPolicy              string          `yaml:"preemptionPolicy"`
Priority                      int             `yaml:"priority"`
RestartPolicy                 string          `yaml:"restartPolicy"`
SchedulerName                 string          `yaml:"schedulerName"`
SecurityContext               SecurityContext `yaml:"securityContext"`
ServiceAccount                string          `yaml:"serviceAccount"`
ServiceAccountName            string          `yaml:"serviceAccountName"`
TerminationGracePeriodSeconds int             `yaml:"terminationGracePeriodSeconds"`
Tolerations                   []Toleration    `yaml:"tolerations"`
Volumes                       []Volumes       `yaml:"volumes"`
}
type Conditions struct {
LastProbeTime      interface{} `yaml:"lastProbeTime"`
LastTransitionTime time.Time   `yaml:"lastTransitionTime"`
Reason             string      `yaml:"reason,omitempty"`
Status             string      `yaml:"status"`
Type               string      `yaml:"type"`
}
type Terminated struct {
ContainerID string    `yaml:"containerID"`
ExitCode    int       `yaml:"exitCode"`
FinishedAt  time.Time `yaml:"finishedAt"`
Reason      string    `yaml:"reason"`
StartedAt   time.Time `yaml:"startedAt"`
}
type State struct {
Terminated Terminated `yaml:"terminated"`
}
type ContainerStatus struct {
ContainerID  string    `yaml:"containerID"`
Image        string    `yaml:"image"`
ImageID   string `yaml:"imageID"`
LastState State  `yaml:"lastState"`
Name      string `yaml:"name"`
Ready        bool      `yaml:"ready"`
RestartCount int       `yaml:"restartCount"`
Started   bool   `yaml:"started"`
State     State  `yaml:"state"`
}
type PodIPs struct {
IP string `yaml:"ip"`
}
type Status struct {
Conditions        []Conditions      `yaml:"conditions"`
ContainerStatuses []ContainerStatus `yaml:"containerStatuses"`
HostIP            string            `yaml:"hostIP"`
Phase             string              `yaml:"phase"`
PodIP             string            `yaml:"podIP"`
PodIPs            []PodIPs          `yaml:"podIPs"`
QosClass          string            `yaml:"qosClass"`
StartTime         time.Time           `yaml:"startTime"`
}
