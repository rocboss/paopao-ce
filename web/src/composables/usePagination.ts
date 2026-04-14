import { ref } from 'vue';

/**
 * 通用分页逻辑 composable
 * 保持与原有代码完全一致的行为
 */
export function usePagination(initialPageSize: number = 20) {
  const loading = ref(false);
  const noMore = ref(false);
  const page = ref(1);
  const pageSize = ref(initialPageSize);
  const totalPage = ref(0);

  const reset = () => {
    loading.value = false;
    noMore.value = false;
    page.value = 1;
    totalPage.value = 0;
  };

  const nextPage = (loadCallback: () => void) => {
    if (page.value < totalPage.value || totalPage.value == 0) {
      noMore.value = false;
      page.value++;
      loadCallback();
    } else {
      noMore.value = true;
    }
  };

  return {
    loading,
    noMore,
    page,
    pageSize,
    totalPage,
    reset,
    nextPage,
  };
}
