/**
 * @file 格式化日期
 */

import moment from "moment";
import "moment/dist/locale/zh-cn";
moment.locale("zh-cn");

export const formatTime = (time: number) => {
  return moment.unix(time).utc(true).format("YYYY-MM-DD HH:mm:ss");
};

export const formatRelativeTime = (time: number) => {
  return moment.unix(time).fromNow();
};

export const formatPrettyTime = (time: number) => {
  let mt = moment.unix(time);
  if (moment().diff(mt, "month") > 3) {
    return mt.utc(true).format("YYYY-MM-DD HH:mm:ss");
  }
  return mt.fromNow();
};
