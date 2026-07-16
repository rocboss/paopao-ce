/**
 * Keep document width stable across:
 * 1) short page → long content (scrollbar appears)
 * 2) n-dialog / n-modal scroll-lock (Naive UI sets overflow:hidden on <html>)
 *
 * Naive UI's useLockHtmlScroll removes the classic scrollbar and optionally
 * adds margin-right compensation. That fights `scrollbar-gutter` / always-on
 * overflow-y:scroll and shifts the flex-centered layout plus fixed sidebars
 * (`left/right: calc(50% + ...)`).
 *
 * Strategy:
 * - CSS keeps the vertical scrollbar track always present (see main.less)
 * - When Naive marks <html> as locked, freeze body scroll with position:fixed
 *   without ever releasing the scrollbar width
 */

let installed = false;

export function installStableScrollbar() {
  if (installed || typeof document === 'undefined') {
    return;
  }
  installed = true;

  const html = document.documentElement;
  let frozen = false;
  let scrollY = 0;

  const isHtmlScrollLocked = () => {
    const { overflow, overflowY } = html.style;
    return overflow === 'hidden' || overflowY === 'hidden';
  };

  const freezeBackgroundScroll = () => {
    if (frozen) {
      return;
    }
    frozen = true;
    scrollY = window.scrollY;
    const body = document.body;
    body.style.position = 'fixed';
    body.style.top = `-${scrollY}px`;
    body.style.left = '0';
    body.style.right = '0';
    body.style.width = '100%';
  };

  const unfreezeBackgroundScroll = () => {
    if (!frozen) {
      return;
    }
    frozen = false;
    const body = document.body;
    body.style.position = '';
    body.style.top = '';
    body.style.left = '';
    body.style.right = '';
    body.style.width = '';
    window.scrollTo(0, scrollY);
  };

  const sync = () => {
    // Drop Naive's margin compensation — scrollbar track already reserves width.
    if (html.style.marginRight && html.style.marginRight !== '0px') {
      html.style.marginRight = '0px';
    }

    if (isHtmlScrollLocked()) {
      freezeBackgroundScroll();
    } else {
      unfreezeBackgroundScroll();
    }
  };

  new MutationObserver(sync).observe(html, {
    attributes: true,
    attributeFilter: ['style'],
  });

  // In case lock styles are already present
  sync();
}
