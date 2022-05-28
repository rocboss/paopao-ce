import { createStore } from 'vuex';

export default createStore({
    state: {
        refresh: Date.now(),
        theme: localStorage.getItem('PAOPAO_THEME'),
        collapsedLeft: document.body.clientWidth <= 821,
        collapsedRight: document.body.clientWidth <= 821,
        authModalShow: false,
        authModelTab: 'signin',
        userInfo: {
            id: 0,
            username: '',
            nickname: ''
        },
    },
    mutations: {
        refresh(state, refresh) {
            state.refresh = refresh || Date.now();
        },
        triggerTheme(state, theme) {
            state.theme = theme;
        },
        triggerAuth(state, status) {
            state.authModalShow = status;
        },
        triggerAuthKey(state, key) {
            state.authModelTab = key;
        },
        triggerCollapsedLeft(state, status) {
            state.collapsedLeft = status;
        },
        triggerCollapsedRight(state, status) {
            state.collapsedRight = status;
        },
        updateUserinfo(state, data) {
            state.userInfo = data;
        },
        userLogout(state) {
            localStorage.removeItem('PAOPAO_TOKEN')
            state.userInfo = { id: 0, nickname: '', username: '' };
        },
    },
    actions: {},
    modules: {},
});
