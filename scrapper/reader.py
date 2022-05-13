from scrapper.setups.yt import youtube

import os
import logging
import scrapper.setups.settings as settings


logger = logging.getLogger(__name__)


def read_file(file_path, separator=settings.TABLE_SEPARATOR):
    if os.path.isfile(file_path):
        with open(file_path, "r") as file:
            data = [line.strip().split(separator) for line in file]
        return data
    else:
        file = open(file_path, "x")
        file.close()
        return []


def get_country_codes(file_path=settings.COUNTRY_CODES_PATH):
    return [line[:2] for line in open(file_path, 'r')]


def write_video_set(video_set, region):
    with open(settings.VIDEO_SET_PATH + region + "/video-set.txt", "x") as file:
        for item in video_set:
            file.write(item + '\n')


def get_video_set(country_code):
    if os.path.isfile(settings.VIDEO_SET_PATH + country_code + "/video-set.txt"):
        return [item.strip() for item in open(settings.VIDEO_SET_PATH + country_code + "/video-set.txt", "r")]
    else:
        video_set = set()

        # for country in codes:
        #     next_page_token = None
        #     while True:
        #         request = youtube[0].videos().list(
        #             part="id, contentDetails",
        #             chart="mostPopular",
        #             regionCode=country_code,
        #             maxResults=50,
        #             pageToken=next_page_token)
        #         response = request.execute()
        #
        #         for item in response.get("items"):
        #             video_set.add(item.get("id"))
        #
        #         next_page_token = response.get('nextPageToken')
        #
        #         if not next_page_token:
        #             break
        next_page_token = None
        while True:
            request = youtube[0].videos().list(
                part="id, contentDetails",
                chart="mostPopular",
                regionCode=country_code,
                maxResults=50,
                pageToken=next_page_token)
            response = request.execute()

            for item in response.get("items"):
                video_set.add(item.get("id"))

            next_page_token = response.get('nextPageToken')

            if not next_page_token:
                break

        video_items = list(video_set)
        write_video_set(video_items, country_code)

        return video_items


class Reader:
    def __init__(self, region):
        self.records = None
        self.region = region
        self.make_structures()

    def make_structures(self):
        logger.debug(f"read records from files")

        views = read_file(settings.VIEWS_PATH + self.region + "/views.tsv")
        likes = read_file(settings.LIKES_PATH + self.region + "/likes.tsv")
        comments = read_file(settings.COMMENTS_PATH + self.region + "/comments.tsv")

        self.records = {
            "views": views,
            "likes": likes,
            "comments": comments,
        }

    def get_records(self):
        return self.records
