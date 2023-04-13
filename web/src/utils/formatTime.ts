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

export const formatPrettyTime = (time: number, noPretty: boolean) => {
  if (noPretty) {
    return moment.unix(time).utc(true).fromNow();
  }
  let mt = moment.unix(time).utc(true);
  let now = moment().utc(true);
  if (mt.year() != now.year()) {
    return mt.format("YYYY-MM-DD HH:mm");
  } else if (moment().diff(mt, "month") > 3) {
    return mt.format("MM-DD HH:mm");
  }
  return mt.fromNow();
};
