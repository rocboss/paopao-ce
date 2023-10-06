export const parsePostTag = (content: string) => {
  const tags: string[] = [];
  const users: string[] = [];
  var tagExp = /(#|＃)([^#@\s])+?\s+?/g; // 这⾥中⽂#和英⽂#都会识别
  var atExp = /@([a-zA-Z0-9])+?\s+?/g; // 这⾥中⽂#和英⽂#都会识别
  content = content
    .replace(/<[^>]*?>/gi, "")
    .replace(/(.*?)<\/[^>]*?>/gi, "")
    .replace(tagExp, (item) => {
      tags.push(item.substr(1).trim());
      return (
        '<a class="hash-link" data-detail="tag:' +
        encodeURIComponent(item.substr(1).trim()) +
        '">' +
        item.trim() +
        "</a> "
      );
    })
    .replace(atExp, (item) => {
      users.push(item.substr(1).trim());
      return (
        '<a class="hash-link" data-detail="user:' +
        encodeURIComponent(item.substr(1).trim()) +
        '">' +
        item.trim() +
        "</a> "
      );
    });
  return { content, tags, users };
};

export const preparePost = (content: string, hint: string, maxSize: number) => {
  let isEllipsis = false;
  if (content.length > maxSize) {
    content = content.substring(0, maxSize);
    isEllipsis = true;
    let latestChar = content.charAt(maxSize - 1);
    if (latestChar == "#" || latestChar == "#" || latestChar == "@") {
      content = content.substring(0, maxSize - 1);
    }
  }
  const tagExp = /(#|＃)([^#@\s])+?\s+?/g; // 这⾥中⽂#和英⽂#都会识别
  const atExp = /@([a-zA-Z0-9])+?\s+?/g; // 这⾥中⽂#和英⽂#都会识别
  content = content
    .replace(/<[^>]*?>/gi, "")
    .replace(/(.*?)<\/[^>]*?>/gi, "")
    .replace(tagExp, (item) => {
      return (
        '<a class="hash-link" data-detail="tag:' +
        encodeURIComponent(item.substring(1).trim()) +
        '">' +
        item.trim() +
        "</a> "
      );
    })
    .replace(atExp, (item) => {
      return (
        '<a class="hash-link" data-detail="user:' +
        encodeURIComponent(item.substring(1).trim()) +
        '">' +
        item.trim() +
        "</a> "
      );
    });
  if (isEllipsis) {
    content =
      content.trimEnd() +
      " ..." +
      '<a class="hash-link" data-detail="post">' +
      hint +
      "</a> ";
  }
  return content;
};
