import datetime
import logging
import scrapper.setups.settings as settings

from shutil import copyfile


logger = logging.getLogger(__name__)


def write_desc_to_file(features, region):
    """
    Write descriptions to file.
    :param features: description features as list.
    :param region: country code of region
    :return: just writing to file without updating any data structures.
    """
    logger.debug(f"adding descriptions")
    with open(settings.DESCRIPTIONS_PATH + region + "/descriptions.tsv", "a") as file:
        file.write("\t".join([str(feature) for feature in features]) + "\n")


class DataHandler:
    def __init__(self, raw_data, records, first_day, region):
        self.raw_data = raw_data
        self.first_day = first_day

        self.region = region

        # Metrics from video history
        self.views = records.get("views")
        self.likes = records.get("likes")
        self.comments = records.get("comments")

        # Run Handler
        logger.debug("Start today handling")
        self.run_handler()

    def run_handler(self):
        if self.first_day:
            self.write_desc()
        self.write_video_stats()

    def write_desc(self):
        for item in self.raw_data:
            video_features = item[0]
            video_id = video_features.get("id")

            logger.debug(f"adding description about video #{video_id}")
            video_desc = [
                # video_id
                video_features.get("id"),
                # title
                video_features.get("snippet").get("title"),
                # published_at
                video_features.get("snippet").get("publishedAt"),
                # category_id
                video_features.get("snippet").get("categoryId"),
                # duration
                video_features.get("contentDetails").get("duration"),
            ]

            logger.debug(f"write descriptions.tsv")
            write_desc_to_file(video_desc, self.region)

    # NOT VALID CHECKING
    def write_video_stats(self):
        metrics = {
            "viewCount": self.views,
            "likeCount": self.likes,
            "commentCount": self.comments
        }

        for key, stat in metrics.items():
            logger.debug(f"updating [{key}] metrics")
            for item in self.raw_data:
                video_features = item[0]
                video_id = video_features.get("id")
                logger.debug(f"checking whether video #{video_id} is new")
                if self.first_day:
                    logger.debug(f"video #{video_id} is new so adding its [{key}] metrics")
                    stat.append([video_id, video_features.get("statistics").get(key)])
                else:
                    logger.debug(f"video #{video_id} is in history so try to append its [{key}] metrics")
                    for video in stat:
                        if video[0] == video_id:
                            logger.debug(f"append metrics because video #{video_id} has no n-days history")
                            video.append(video_features.get("statistics").get(key))
                            break

    def save_data(self):
        # YYYY-MM-DD:
        cur_period = "-" + str(datetime.datetime.now().isoformat()[:10]) + settings.TABLE_FORMAT
        # full date:
        # cur_period = "-" + str(datetime.datetime.now().isoformat()) + settings.TABLE_FORMAT

        path_data_pairs = [
            (settings.VIEWS_PATH + self.region + "/views.tsv", self.views),
            (settings.LIKES_PATH + self.region + "/likes.tsv", self.likes),
            (settings.COMMENTS_PATH + self.region + "/comments.tsv", self.comments),
        ]

        for path, data in path_data_pairs:
            file_path = path[:-4] + cur_period
            with open(file_path, "x") as file:
                for item in data:
                    file.write(settings.TABLE_SEPARATOR.join([str(elem) for elem in item]) + "\n")
            copyfile(file_path, path)
