package domain

type Face struct {
	Confidence  float64     `json:"confidence"`
	Coordinates Coordinates `json:"coordinates"`
	FaceID      string      `json:"face_id"`
}

type Coordinates struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Xmax   int `json:"xmax"`
	Xmin   int `json:"xmin"`
	Ymax   int `json:"ymax"`
	Ymin   int `json:"ymin"`
}

type FaceDetectionResult struct {
	Faces []Face `json:"faces"`
}

type FaceSimilarityResult struct {
	Score float64 `json:"score"`
}
