export const prettyQuoteNum = (num: number) => {
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + '千';
  } else if (num >= 10000) {
    return (num / 10000).toFixed(1) + '万';
  }
  return num;
};
