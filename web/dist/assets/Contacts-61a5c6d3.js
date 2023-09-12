import{_ as T}from"./whisper-24717711.js";import{d as N,c as j,r as A,e as s,f as c,k as t,w as n,j as i,y as H,A as L,x as v,bf as g,h as I,H as a,b as U,Y as S,F as z,u as W,q as E}from"./@vue-a481fc63.js";import{b as G}from"./formatTime-4210fcd1.js";import{i as Q,p as Y}from"./@vicons-7a4ef312.js";import{j as x,o as J,e as K,O as X,L as Z,F as ee,Q as te,I as ne,G as oe}from"./naive-ui-d8de3dda.js";import{_ as q,b as se}from"./index-d9671de1.js";import{_ as ae}from"./post-skeleton-73b6be04.js";import{_ as ce}from"./main-nav.vue_vue_type_style_index_0_lang-5497f713.js";import{u as ie}from"./vuex-44de225f.js";import{b as _e}from"./vue-router-e5a2430e.js";import"./moment-2ab8298d.js";import"./seemly-76b7b838.js";import"./vueuc-39372edb.js";import"./evtd-b614532e.js";import"./@css-render-7124a1a5.js";import"./vooks-6d99783e.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";import"./axios-4a70c6fc.js";/* empty css               */const re={class:"contact-item"},le={class:"nickname-wrap"},pe={class:"username-wrap"},ue={class:"user-info"},me={class:"info-item"},de={class:"info-item"},fe={class:"item-header-extra"},ve=N({__name:"contact-item",props:{contact:{}},emits:["send-whisper"],setup(b,{emit:h}){const _=b,r=e=>()=>I(x,null,{default:()=>I(e)}),l=j(()=>[{label:"私信",key:"whisper",icon:r(Y)}]),u=e=>{switch(e){case"whisper":const o={id:_.contact.user_id,avatar:_.contact.avatar,username:_.contact.username,nickname:_.contact.nickname,is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1};h("send-whisper",o);break}};return(e,o)=>{const m=J,d=A("router-link"),w=K,k=X,y=Z;return s(),c("div",re,[t(y,{"content-indented":""},{avatar:n(()=>[t(m,{size:54,src:e.contact.avatar},null,8,["src"])]),header:n(()=>[i("span",le,[t(d,{onClick:o[0]||(o[0]=H(()=>{},["stop"])),class:"username-link",to:{name:"user",query:{s:e.contact.username}}},{default:n(()=>[L(v(e.contact.nickname),1)]),_:1},8,["to"])]),i("span",pe," @"+v(e.contact.username),1),i("div",ue,[i("span",me," UID. "+v(e.contact.user_id),1),i("span",de,v(g(G)(e.contact.created_on))+" 加入 ",1)])]),"header-extra":n(()=>[i("div",fe,[t(k,{placement:"bottom-end",trigger:"click",size:"small",options:l.value,onSelect:u},{default:n(()=>[t(w,{quaternary:"",circle:""},{icon:n(()=>[t(g(x),null,{default:n(()=>[t(g(Q))]),_:1})]),_:1})]),_:1},8,["options"])])]),_:1})])}}});const ge=q(ve,[["__scopeId","data-v-d62f19da"]]),he={key:0,class:"skeleton-wrap"},we={key:1},ke={key:0,class:"empty-wrap"},ye={key:0,class:"pagination-wrap"},be=N({__name:"Contacts",setup(b){const h=ie(),_=_e(),r=a(!1),l=a([]),u=a(+_.query.p||1),e=a(20),o=a(0),m=a(!1),d=a({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1}),w=p=>{d.value=p,m.value=!0},k=()=>{m.value=!1},y=p=>{u.value=p,C()};U(()=>{C()});const C=(p=!1)=>{l.value.length===0&&(r.value=!0),se({page:u.value,page_size:e.value}).then(f=>{r.value=!1,l.value=f.list,o.value=Math.ceil(f.pager.total_rows/e.value),p&&setTimeout(()=>{window.scrollTo(0,99999)},50)}).catch(f=>{r.value=!1})};return(p,f)=>{const B=ce,F=ae,M=ne,P=ge,V=oe,D=T,O=ee,R=te;return s(),c(z,null,[i("div",null,[t(B,{title:"好友"}),t(O,{class:"main-content-wrap",bordered:""},{default:n(()=>[r.value?(s(),c("div",he,[t(F,{num:e.value},null,8,["num"])])):(s(),c("div",we,[l.value.length===0?(s(),c("div",ke,[t(M,{size:"large",description:"暂无数据"})])):S("",!0),(s(!0),c(z,null,W(l.value,$=>(s(),E(V,{class:"list-item",key:$.user_id},{default:n(()=>[t(P,{contact:$,onSendWhisper:w},null,8,["contact"])]),_:2},1024))),128))])),t(D,{show:m.value,user:d.value,onSuccess:k},null,8,["show","user"])]),_:1})]),o.value>0?(s(),c("div",ye,[t(R,{page:u.value,"onUpdate:page":y,"page-slot":g(h).state.collapsedRight?5:8,"page-count":o.value},null,8,["page","page-slot","page-count"])])):S("",!0)],64)}}});const Ye=q(be,[["__scopeId","data-v-e20fef94"]]);export{Ye as default};
