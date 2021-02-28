package main

type SonarqubeQualityGateCondition struct {
	ErrorThreshold string `json:"errorThreshold"`
	Metric string `json:"metric"`
	Status string `json:"status"`
	Operator string `json:"operator"`
	Value string `json:"value"`
}

type SonarqubeQualityGate struct {
	Status string `json:"status"`
	Conditions []SonarqubeQualityGateCondition
}

type SonarqubeProject struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
}

type SonarqubeMessage struct {
	ServerUrl string `json:"serverUrl"`
	TaskId string `json:"taskId"`
	Status string `json:"status"`
	AnalyzeDate string `json:"analysedAt"`
	Project SonarqubeProject `json:"project"`
	QualityGade SonarqubeQualityGate `json:"qualityGate"`
}
