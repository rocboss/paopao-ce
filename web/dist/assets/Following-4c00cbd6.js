import{_ as Y}from"./whisper-7c8226c3.js";import{d as B,c as K,r as Q,e as _,f as u,k as o,w as t,j as p,y as X,A as P,x as k,q as N,Y as U,bf as g,h as S,H as r,b as Z,F as C,u as ee}from"./@vue-a481fc63.js";import{u as oe,b as ne}from"./vue-router-e5a2430e.js";import{I as te,u as se,f as ae,_ as O,V as le,W as ce}from"./index-537a241d.js";import{k as ie,r as _e,G as re,x as ue}from"./@vicons-4ed09704.js";import{F as pe,j as M,o as me,M as de,e as fe,P as ge,O as we,G as ve,R as he,f as ke,g as ye,J as be,H as $e}from"./naive-ui-defd0b2d.js";import{_ as Fe}from"./post-skeleton-5d89319a.js";import{_ as ze}from"./main-nav.vue_vue_type_style_index_0_lang-832ff0ed.js";import{u as Te}from"./vuex-44de225f.js";import"./axios-4a70c6fc.js";import"./moment-2ab8298d.js";/* empty css               */import"./seemly-76b7b838.js";import"./vueuc-39372edb.js";import"./evtd-b614532e.js";import"./@css-render-7124a1a5.js";import"./vooks-6d99783e.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";const Ue={class:"follow-item"},qe={class:"nickname-wrap"},xe={class:"username-wrap"},Ie={class:"user-info"},Pe={class:"info-item"},Se={class:"info-item"},Ce={class:"item-header-extra"},Me=B({__name:"follow-item",props:{contact:{}},emits:["send-whisper"],setup(q,{emit:y}){const n=q,c=pe();oe();const s=e=>()=>S(M,null,{default:()=>S(e)}),b=()=>{c.success({title:"提示",content:"确定"+(n.contact.is_following?"取消关注":"关注")+"该用户吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{n.contact.is_following?se({user_id:n.contact.user_id}).then(e=>{window.$message.success("取消关注成功"),n.contact.is_following=!1}).catch(e=>{console.log(e)}):ae({user_id:n.contact.user_id}).then(e=>{window.$message.success("关注成功"),n.contact.is_following=!0}).catch(e=>{console.log(e)})}})},w=K(()=>{let e=[{label:"私信",key:"whisper",icon:s(_e)}];return n.contact.is_following?e.push({label:"取消关注",key:"unfollow",icon:s(re)}):e.push({label:"关注",key:"follow",icon:s(ue)}),e}),m=e=>{switch(e){case"follow":case"unfollow":b();break;case"whisper":const a={id:n.contact.user_id,avatar:n.contact.avatar,username:n.contact.username,nickname:n.contact.nickname,is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1};y("send-whisper",a);break}};return(e,a)=>{const d=me,f=Q("router-link"),v=de,$=fe,F=ge,z=we;return _(),u("div",Ue,[o(z,{"content-indented":""},{avatar:t(()=>[o(d,{size:54,src:e.contact.avatar},null,8,["src"])]),header:t(()=>[p("span",qe,[o(f,{onClick:a[0]||(a[0]=X(()=>{},["stop"])),class:"username-link",to:{name:"user",query:{s:e.contact.username}}},{default:t(()=>[P(k(e.contact.nickname),1)]),_:1},8,["to"])]),p("span",xe," @"+k(e.contact.username),1),e.contact.is_following?(_(),N(v,{key:0,class:"top-tag",type:"success",size:"small",round:""},{default:t(()=>[P(" 已关注 ")]),_:1})):U("",!0),p("div",Ie,[p("span",Pe," UID. "+k(e.contact.user_id),1),p("span",Se,k(g(te)(e.contact.created_on))+" 加入 ",1)])]),"header-extra":t(()=>[p("div",Ce,[o(F,{placement:"bottom-end",trigger:"click",size:"small",options:w.value,onSelect:m},{default:t(()=>[o($,{quaternary:"",circle:""},{icon:t(()=>[o(g(M),null,{default:t(()=>[o(g(ie))]),_:1})]),_:1})]),_:1},8,["options"])])]),_:1})])}}});const Be=O(Me,[["__scopeId","data-v-1fb7364a"]]),Ne={key:0,class:"skeleton-wrap"},Oe={key:1},Re={key:0,class:"empty-wrap"},Ve={key:0,class:"pagination-wrap"},De=B({__name:"Following",setup(q){const y=Te(),n=ne(),c=r(!1),s=r([]),b=n.query.n||"粉丝详情",w=n.query.s||"",m=r(n.query.t||"follows"),e=r(+n.query.p||1),a=r(20),d=r(0),f=r(!1),v=r({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1}),$=l=>{v.value=l,f.value=!0},F=()=>{f.value=!1},z=l=>{e.value=l,T()},R=l=>{m.value=l,T()},T=()=>{m.value==="follows"?V(w):m.value==="followings"&&D(w)},V=(l,h=!1)=>{s.value.length===0&&(c.value=!0),le({username:l,page:e.value,page_size:a.value}).then(i=>{c.value=!1,s.value=i.list||[],d.value=Math.ceil(i.pager.total_rows/a.value),h&&setTimeout(()=>{window.scrollTo(0,99999)},50)}).catch(i=>{c.value=!1})},D=(l,h=!1)=>{s.value.length===0&&(c.value=!0),ce({username:l,page:e.value,page_size:a.value}).then(i=>{c.value=!1,s.value=i.list||[],d.value=Math.ceil(i.pager.total_rows/a.value),h&&setTimeout(()=>{window.scrollTo(0,99999)},50)}).catch(i=>{c.value=!1})};return Z(()=>{T()}),(l,h)=>{const i=ze,x=ke,W=ye,H=Fe,j=be,A=Be,G=$e,E=Y,J=ve,L=he;return _(),u(C,null,[p("div",null,[o(i,{title:g(b),back:!0},null,8,["title"]),o(J,{class:"main-content-wrap",bordered:""},{default:t(()=>[o(W,{type:"line",animated:"","default-value":m.value,"onUpdate:value":R},{default:t(()=>[o(x,{name:"follows",tab:"正在关注"}),o(x,{name:"followings",tab:"我的粉丝"})]),_:1},8,["default-value"]),c.value?(_(),u("div",Ne,[o(H,{num:a.value},null,8,["num"])])):(_(),u("div",Oe,[s.value.length===0?(_(),u("div",Re,[o(j,{size:"large",description:"暂无数据"})])):U("",!0),(_(!0),u(C,null,ee(s.value,I=>(_(),N(G,{key:I.user_id},{default:t(()=>[o(A,{contact:I,onSendWhisper:$},null,8,["contact"])]),_:2},1024))),128))])),o(E,{show:f.value,user:v.value,onSuccess:F},null,8,["show","user"])]),_:1})]),d.value>0?(_(),u("div",Ve,[o(L,{page:e.value,"onUpdate:page":z,"page-slot":g(y).state.collapsedRight?5:8,"page-count":d.value},null,8,["page","page-slot","page-count"])])):U("",!0)],64)}}});const po=O(De,[["__scopeId","data-v-0a10234f"]]);export{po as default};
