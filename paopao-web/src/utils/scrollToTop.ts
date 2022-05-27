// 滚动到顶部
export const scrollToTop = (scrollDuration: number) => {
    var cosParameter = window.scrollY / 2;
    var scrollCount = 0;
    var oldTimestamp = performance.now();
    function step(newTimestamp: number) {
        scrollCount +=
            Math.PI / (scrollDuration / (newTimestamp - oldTimestamp));
        if (scrollCount >= Math.PI) window.scrollTo(0, 0);
        if (window.scrollY === 0) return;
        window.scrollTo(
            0,
            Math.round(cosParameter + cosParameter * Math.cos(scrollCount))
        );
        oldTimestamp = newTimestamp;
        window.requestAnimationFrame(step);
    }
    window.requestAnimationFrame(step);
};