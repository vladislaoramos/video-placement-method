import csv
import pandas as pd


DESCRIPTION_FEATURES = [
    "video_id", 
    "title", 
    "published_at", 
    "category_id", 
    "duration",
]

DAYS = 8


def get_df(path, cols):
    with open(path, "r") as file:
        return pd.DataFrame(csv.reader(file, delimiter="\t"), columns=cols)


def get_daily_cols(metrics):
    return ["video_id"] + [metrics + '_' + str(i) for i in range(1, DAYS + 1)]

regions = [
    "FI", "KZ", "UA", "EG", "IT", "PT",
    "GB", "BE", "NO", "LV", "RU", "TR",
    "ES", "FR", "SE", "PL", "NL", "DE"]


def main():
    for region in regions:
        desc_df = get_df("features/descriptions/" + region + "/descriptions.tsv", DESCRIPTION_FEATURES)
        
        views_df = get_df("features/views/" + region + "/views.tsv", get_daily_cols("views"))
        likes_df = get_df("features/likes/" + region + "/likes.tsv", get_daily_cols("likes"))
        comments_df = get_df("features/comments/" + region + "/comments.tsv", get_daily_cols("comments"))

        dynamic_features = views_df.merge(likes_df, on="video_id").merge(comments_df, on="video_id")

        # make whole dataset
        data = desc_df.merge(dynamic_features, on="video_id")
        data.to_csv("data_" + region + ".tsv", sep='\t', index=False)


if __name__ == "__main__":
    main()