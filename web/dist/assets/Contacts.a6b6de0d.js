import{ah as k,d as C,aj as D,V as e,W as n,Y as c,a3 as o,a9 as v,ad as E,Z as N,ai as P,r as l,a1 as R,cn as S,a4 as g,a2 as U,a6 as y,aa as q,ab as M,a5 as T}from"./index.f480f018.js";import{_ as j}from"./post-skeleton.39a35f5b.js";import{_ as F}from"./main-nav.9110259b.js";import{_ as L}from"./List.49bcdf81.js";import{_ as W}from"./Pagination.1263479e.js";import{a as Y,_ as Z}from"./Skeleton.57e98fe3.js";const A={class:"avatar"},G={class:"base-info"},H={class:"username"},J={class:"uid"},K=C({props:{contact:null},setup(s){const p=D(),m=t=>{p.push({name:"user",query:{username:t}})};return(t,a)=>{const _=E;return e(),n("div",{class:"contact-item",onClick:a[0]||(a[0]=u=>m(s.contact.username))},[c("div",A,[o(_,{size:"large",src:s.contact.avatar},null,8,["src"])]),c("div",G,[c("div",H,[c("strong",null,v(s.contact.nickname),1),c("span",null," @"+v(s.contact.username),1)]),c("div",J,"UID. "+v(s.contact.user_id),1)])])}}});var O=k(K,[["__scopeId","data-v-23bc18c8"]]);const Q={key:0,class:"pagination-wrap"},X={key:0,class:"skeleton-wrap"},ee={key:1},te={key:0,class:"empty-wrap"},ae=C({setup(s){const p=N(),m=P(),t=l(!1),a=l([]),_=l(+m.query.p||1),u=l(20),d=l(0),w=r=>{_.value=r,f()};R(()=>{f()});const f=(r=!1)=>{a.value.length===0&&(t.value=!0),S({page:_.value,page_size:u.value}).then(i=>{t.value=!1,a.value=i.list,d.value=Math.ceil(i.pager.total_rows/u.value),r&&setTimeout(()=>{window.scrollTo(0,99999)},50)}).catch(i=>{t.value=!1})};return(r,i)=>{const x=F,$=W,B=j,b=Y,z=O,I=Z,V=L;return e(),n("div",null,[o(x,{title:"\u597D\u53CB"}),o(V,{class:"main-content-wrap",bordered:""},{footer:g(()=>[d.value>1?(e(),n("div",Q,[o($,{page:_.value,"onUpdate:page":w,"page-slot":U(p).state.collapsedRight?5:8,"page-count":d.value},null,8,["page","page-slot","page-count"])])):y("",!0)]),default:g(()=>[t.value?(e(),n("div",X,[o(B,{num:u.value},null,8,["num"])])):(e(),n("div",ee,[a.value.length===0?(e(),n("div",te,[o(b,{size:"large",description:"\u6682\u65E0\u6570\u636E"})])):y("",!0),(e(!0),n(q,null,M(a.value,h=>(e(),T(I,{key:h.user_id},{default:g(()=>[o(z,{contact:h},null,8,["contact"])]),_:2},1024))),128))]))]),_:1})])}}});var ue=k(ae,[["__scopeId","data-v-e9de9e50"]]);export{ue as default};