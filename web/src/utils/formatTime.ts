/**
 * @file 格式化日期
 */

import moment from "moment";
import "moment/dist/locale/zh-cn";
moment.locale("zh-cn");

export const formatTime = (time: number) => {
  return moment.unix(time).utc(true).format("YYYY-MM-DD HH:mm");
};

export const formatRelativeTime = (time: number) => {
  return moment.unix(time).fromNow();
};

export const formatPrettyTime = (time: number) => {
  let mt = moment.unix(time);
  let now = moment();
  if (mt.year() != now.year()) {
    return mt.utc(true).format("YYYY-MM-DD HH:mm");
  } else if (moment().diff(mt, "month") > 3) {
    return mt.utc(true).format("MM-DD HH:mm");
  }
  return mt.fromNow();
};

export const formatPrettyDate = (time: number) => {
  let mt = moment.unix(time);
  let now = moment();
  if (mt.year() != now.year()) {
    return mt.utc(true).format("YYYY-MM-DD");
  } else if (moment().diff(mt, "month") > 3) {
    return mt.utc(true).format("MM-DD");
  }
  return mt.fromNow();
};

export const formatDate = (time: number) => {
  return moment.unix(time).utc(true).format("YYYY年MM月");
};
