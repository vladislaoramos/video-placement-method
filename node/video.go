package node

type VideoObject struct {
	Id       string  `json:"video_id"`
	Title    string  `json:"title"`
	Size     float64 `json:"size"`
	Category float64 `json:"category"`

	ViewsLastDay    float64 `json:"views"`
	LikesLastDay    float64 `json:"likes"`
	CommentsLastDay float64 `json:"comments"`

	ViewsWMA    float64 `json:"views_wma"`
	LikesWMA    float64 `json:"likes_wma"`
	CommentsWMA float64 `json:"comments_wma"`

	// calc
	ViewsLikesRatio        float64 `json:"views-likes-ratio"`
	ViewsLikesCommRatio    float64 `json:"views_likes_comm_ratio"`
	ViewsLikesCommCatRatio float64 `json:"views_likes_comm_cat_ratio"`
	ViewsCommRatio         float64 `json:"views_comm_ratio"`
	ViewsCommCatRatio      float64 `json:"views_comm_cat_ratio"`
	ViewsCatRatio          float64 `json:"views_cat_ratio"`
	ViewsLikesCatRatio     float64 `json:"views_likes_cat_ratio"`
}

type VideoList []VideoObject

func (v *VideoObject) setViewsLikesRatio(weights []float64) {
	//v.ViewsLikesRatio = 0.8*v.ViewsLastDay + 0.2*v.LikesLastDay
	v.ViewsLikesRatio = weights[0]*v.ViewsLastDay + weights[1]*v.LikesLastDay
}

func (v *VideoObject) setViewsLikesWMA(weights []float64) {
	//v.ViewsLikesRatio = 0.8*v.ViewsWMA + 0.2*v.LikesWMA
	v.ViewsLikesRatio = weights[0]*v.ViewsWMA + weights[1]*v.LikesWMA
}

func (v *VideoObject) setViewsLikesCommRatio(weights []float64) {
	//v.ViewsLikesCommRatio = 0.6*v.ViewsLastDay + 0.3*v.CommentsLastDay + 0.1*v.LikesLastDay
	v.ViewsLikesCommRatio = weights[0]*v.ViewsLastDay + weights[2]*v.CommentsLastDay + weights[1]*v.LikesLastDay
}

func (v *VideoObject) setViewsLikesCommWMA(weights []float64) {
	//v.ViewsLikesCommRatio = 0.6*v.ViewsWMA + 0.3*v.CommentsWMA + 0.1*v.LikesWMA
	v.ViewsLikesCommRatio = weights[0]*v.ViewsWMA + weights[2]*v.CommentsWMA + weights[1]*v.LikesWMA
}

func (v *VideoObject) setViewsLikesCommCatRatio(weights []float64) {
	//v.ViewsLikesCommCatRatio =
	//	0.6*v.ViewsLastDay + 0.2*v.CommentsLastDay + 0.1*v.LikesLastDay + 0.1*v.Category
	v.ViewsLikesCommCatRatio =
		weights[0]*v.ViewsLastDay + weights[2]*v.CommentsLastDay + weights[1]*v.LikesLastDay + weights[3]*v.Category
}

func (v *VideoObject) setViewsLikesCommCatWMA(weights []float64) {
	//v.ViewsLikesCommCatRatio =
	//	0.6*v.ViewsWMA + 0.2*v.CommentsWMA + 0.1*v.LikesWMA + 0.1*v.Category
	v.ViewsLikesCommCatRatio =
		weights[0]*v.ViewsWMA + weights[2]*v.CommentsWMA + weights[1]*v.LikesWMA + weights[3]*v.Category
}

//

func (v *VideoObject) setViewsCommRatio(weights []float64) {
	//v.ViewsCommRatio = 0.7*v.ViewsLastDay + 0.3*v.CommentsLastDay
	v.ViewsCommRatio = weights[0]*v.ViewsLastDay + weights[2]*v.CommentsLastDay
}

func (v *VideoObject) setViewsCommWMA(weights []float64) {
	//v.ViewsCommRatio = 0.7*v.ViewsWMA + 0.3*v.CommentsWMA
	v.ViewsCommRatio = weights[0]*v.ViewsWMA + weights[2]*v.CommentsWMA
}

func (v *VideoObject) setViewsCommCatRatio(weights []float64) {
	//v.ViewsCommCatRatio = 0.7*v.ViewsLastDay + 0.2*v.CommentsLastDay + 0.1*v.Category
	v.ViewsCommCatRatio = weights[0]*v.ViewsLastDay + weights[2]*v.CommentsLastDay + weights[3]*v.Category
}

func (v *VideoObject) setViewsCommCatWMA(weights []float64) {
	//v.ViewsCommCatRatio = 0.7*v.ViewsWMA + 0.2*v.CommentsWMA + 0.1*v.Category
	v.ViewsCommCatRatio = weights[0]*v.ViewsWMA + weights[2]*v.CommentsWMA + weights[3]*v.Category
}

func (v *VideoObject) setViewsCatRatio(weights []float64) {
	//v.ViewsCatRatio = 0.8*v.ViewsLastDay + 0.2*v.Category
	v.ViewsCatRatio = weights[0]*v.ViewsLastDay + weights[3]*v.Category
}

func (v *VideoObject) setViewsCatWMA(weights []float64) {
	//v.ViewsCatRatio = 0.8*v.ViewsWMA + 0.2*v.Category
	v.ViewsCatRatio = weights[0]*v.ViewsWMA + weights[3]*v.Category
}

func (v *VideoObject) setViewsLikesCatRatio(weights []float64) {
	//v.ViewsLikesCatRatio = 0.7*v.ViewsLastDay + 0.1*v.LikesLastDay + 0.1*v.Category
	v.ViewsLikesCatRatio = weights[0]*v.ViewsLastDay + weights[1]*v.LikesLastDay + weights[3]*v.Category
}

func (v *VideoObject) setViewsLikesCatWMA(weights []float64) {
	//v.ViewsLikesCatRatio = 0.7*v.ViewsWMA + 0.1*v.LikesWMA*0.1*v.Category
	v.ViewsLikesCatRatio = weights[0]*v.ViewsWMA + weights[1]*v.LikesWMA + weights[3]*v.Category
}
