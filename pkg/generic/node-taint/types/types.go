package types

import (
	clientTypes "k8s.io/apimachinery/pkg/types"
)

// ExperimentDetails is for collecting all the experiment-related details
type ExperimentDetails struct {
	ExperimentName     string
	EngineName         string
	RampTime           int
	ChaosDuration      int
	ChaosLib           string
	AppNS              string
	AppLabel           string
	AppKind            string
	ChaosUID           clientTypes.UID
	InstanceID         string
	ChaosNamespace     string
	ChaosPodName       string
	TargetNode         string
	AuxiliaryAppInfo   string
	Taints             string
	Timeout            int
	Delay              int
	LIBImagePullPolicy string
	TargetContainer    string
	NodeLabel          string
	SetHelperData      string
}
