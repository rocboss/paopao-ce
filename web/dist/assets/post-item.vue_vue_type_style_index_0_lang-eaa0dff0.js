import{p as R,a as V,_ as E,b as Y,c as G}from"./content-1a1bcb51.js";import{d as J,c as x,r as K,e as a,f,k as n,al as Q,w as o,j as i,F as U,u as W,y as m,bf as l,A as r,x as p,q as u,Y as c,h as C}from"./@vue-a481fc63.js";import{u as X}from"./vuex-44de225f.js";import{u as Z}from"./vue-router-e5a2430e.js";import{R as ee,w as te,x as se}from"./index-fae12ace.js";import{c as oe}from"./copy-to-clipboard-4ef7d3eb.js";import{i as ne,j as ae,l as le,m as ie,p as ue,o as ce}from"./@vicons-7a4ef312.js";import{j as v,o as re,M as pe,e as _e,O as me,a as ve,L as de}from"./naive-ui-d8de3dda.js";const he={class:"post-item"},ge={class:"nickname-wrap"},ke={class:"username-wrap"},ye={class:"timestamp-mobile"},fe={class:"item-header-extra"},we=["innerHTML"],be=["onClick"],xe=["onClick"],Ne=J({__name:"mobile-post-item",props:{post:{}},emits:["send-whisper"],setup($,{emit:z}){const d=$,g=Z(),T=X(),w=t=>()=>C(v,null,{default:()=>C(t)}),q=x(()=>{let t=[];return t.push({label:"私信",key:"whisper",icon:w(ue)}),t.push({label:"复制链接",key:"copyTweetLink",icon:w(ce)}),t}),P=async t=>{switch(t){case"copyTweetLink":oe(`${window.location.origin}/#/post?id=${e.value.id}&share=copy_link&t=${new Date().getTime()}`),window.$message.success("链接已复制到剪贴板");break;case"whisper":z("send-whisper",d.post.user);break}},e=x({get:()=>{let t=Object.assign({texts:[],imgs:[],videos:[],links:[],attachments:[],charge_attachments:[]},d.post);return t.contents.map(s=>{(+s.type==1||+s.type==2)&&t.texts.push(s),+s.type==3&&t.imgs.push(s),+s.type==4&&t.videos.push(s),+s.type==6&&t.links.push(s),+s.type==7&&t.attachments.push(s),+s.type==8&&t.charge_attachments.push(s)}),t},set:t=>{d.post.upvote_count=t.upvote_count,d.post.collection_count=t.collection_count}}),L=()=>{te({id:e.value.id}).then(t=>{t.status?e.value={...e.value,upvote_count:e.value.upvote_count+1}:e.value={...e.value,upvote_count:e.value.upvote_count>0?e.value.upvote_count-1:0}}).catch(t=>{console.log(t)})},O=()=>{se({id:e.value.id}).then(t=>{t.status?e.value={...e.value,collection_count:e.value.collection_count+1}:e.value={...e.value,collection_count:e.value.collection_count>0?e.value.collection_count-1:0}}).catch(t=>{console.log(t)})},k=t=>{g.push({name:"post",query:{id:t}})},S=(t,s)=>{if(t.target.dataset.detail){const _=t.target.dataset.detail.split(":");if(_.length===2){T.commit("refresh"),_[0]==="tag"?g.push({name:"home",query:{q:_[1],t:"tag"}}):g.push({name:"user",query:{s:_[1]}});return}}k(s)};return(t,s)=>{const _=re,j=K("router-link"),y=pe,M=_e,D=me,b=V,H=E,B=Y,I=G,N=ve,A=de;return a(),f("div",he,[n(A,{"content-indented":""},Q({avatar:o(()=>[n(_,{round:"",size:30,src:e.value.user.avatar},null,8,["src"])]),header:o(()=>[i("span",ge,[n(j,{onClick:s[0]||(s[0]=m(()=>{},["stop"])),class:"username-link",to:{name:"user",query:{s:e.value.user.username}}},{default:o(()=>[r(p(e.value.user.nickname),1)]),_:1},8,["to"])]),i("span",ke," @"+p(e.value.user.username),1),e.value.is_top?(a(),u(y,{key:0,class:"top-tag",type:"warning",size:"small",round:""},{default:o(()=>[r(" 置顶 ")]),_:1})):c("",!0),e.value.visibility==1?(a(),u(y,{key:1,class:"top-tag",type:"error",size:"small",round:""},{default:o(()=>[r(" 私密 ")]),_:1})):c("",!0),e.value.visibility==2?(a(),u(y,{key:2,class:"top-tag",type:"info",size:"small",round:""},{default:o(()=>[r(" 好友可见 ")]),_:1})):c("",!0),i("div",null,[i("span",ye,p(l(ee)(e.value.created_on))+" "+p(e.value.ip_loc),1)])]),"header-extra":o(()=>[i("div",fe,[n(D,{placement:"bottom-end",trigger:"click",size:"small",options:q.value,onSelect:P},{default:o(()=>[n(M,{quaternary:"",circle:""},{icon:o(()=>[n(l(v),null,{default:o(()=>[n(l(ne))]),_:1})]),_:1})]),_:1},8,["options"])])]),footer:o(()=>[e.value.attachments.length>0?(a(),u(b,{key:0,attachments:e.value.attachments},null,8,["attachments"])):c("",!0),e.value.charge_attachments.length>0?(a(),u(b,{key:1,attachments:e.value.charge_attachments,price:e.value.attachment_price},null,8,["attachments","price"])):c("",!0),e.value.imgs.length>0?(a(),u(H,{key:2,imgs:e.value.imgs},null,8,["imgs"])):c("",!0),e.value.videos.length>0?(a(),u(B,{key:3,videos:e.value.videos},null,8,["videos"])):c("",!0),e.value.links.length>0?(a(),u(I,{key:4,links:e.value.links},null,8,["links"])):c("",!0)]),action:o(()=>[n(N,{justify:"space-between"},{default:o(()=>[i("div",{class:"opt-item",onClick:m(L,["stop"])},[n(l(v),{size:"18",class:"opt-item-icon"},{default:o(()=>[n(l(ae))]),_:1}),r(" "+p(e.value.upvote_count),1)],8,be),i("div",{class:"opt-item",onClick:s[3]||(s[3]=m(h=>k(e.value.id),["stop"]))},[n(l(v),{size:"18",class:"opt-item-icon"},{default:o(()=>[n(l(le))]),_:1}),r(" "+p(e.value.comment_count),1)]),i("div",{class:"opt-item",onClick:m(O,["stop"])},[n(l(v),{size:"18",class:"opt-item-icon"},{default:o(()=>[n(l(ie))]),_:1}),r(" "+p(e.value.collection_count),1)],8,xe)]),_:1})]),_:2},[e.value.texts.length>0?{name:"description",fn:o(()=>[i("div",{onClick:s[2]||(s[2]=h=>k(e.value.id))},[(a(!0),f(U,null,W(e.value.texts,h=>(a(),f("span",{key:h.id,class:"post-text",onClick:s[1]||(s[1]=m(F=>S(F,e.value.id),["stop"])),innerHTML:l(R)(h.content).content},null,8,we))),128))])]),key:"0"}:void 0]),1024)])}}});const Ce={class:"nickname-wrap"},$e={class:"username-wrap"},ze={class:"item-header-extra"},Te={class:"timestamp"},qe=["innerHTML"],Pe=["onClick"],Le=["onClick"],Ae=J({__name:"post-item",props:{post:{}},emits:["send-whisper"],setup($,{emit:z}){const d=$,g=Z(),T=X(),w=t=>()=>C(v,null,{default:()=>C(t)}),q=x(()=>{let t=[];return t.push({label:"私信",key:"whisper",icon:w(ue)}),t.push({label:"复制链接",key:"copyTweetLink",icon:w(ce)}),t}),P=async t=>{switch(t){case"copyTweetLink":oe(`${window.location.origin}/#/post?id=${e.value.id}&share=copy_link&t=${new Date().getTime()}`),window.$message.success("链接已复制到剪贴板");break;case"whisper":z("send-whisper",d.post.user);break}},e=x({get:()=>{let t=Object.assign({texts:[],imgs:[],videos:[],links:[],attachments:[],charge_attachments:[]},d.post);return t.contents.map(s=>{(+s.type==1||+s.type==2)&&t.texts.push(s),+s.type==3&&t.imgs.push(s),+s.type==4&&t.videos.push(s),+s.type==6&&t.links.push(s),+s.type==7&&t.attachments.push(s),+s.type==8&&t.charge_attachments.push(s)}),t},set:t=>{d.post.upvote_count=t.upvote_count,d.post.collection_count=t.collection_count}}),L=()=>{te({id:e.value.id}).then(t=>{t.status?e.value={...e.value,upvote_count:e.value.upvote_count+1}:e.value={...e.value,upvote_count:e.value.upvote_count>0?e.value.upvote_count-1:0}}).catch(t=>{console.log(t)})},O=()=>{se({id:e.value.id}).then(t=>{t.status?e.value={...e.value,collection_count:e.value.collection_count+1}:e.value={...e.value,collection_count:e.value.collection_count>0?e.value.collection_count-1:0}}).catch(t=>{console.log(t)})},k=t=>{g.push({name:"post",query:{id:t}})},S=(t,s)=>{if(t.target.dataset.detail){const _=t.target.dataset.detail.split(":");if(_.length===2){T.commit("refresh"),_[0]==="tag"?g.push({name:"home",query:{q:_[1],t:"tag"}}):g.push({name:"user",query:{s:_[1]}});return}}k(s)};return(t,s)=>{const _=re,j=K("router-link"),y=pe,M=_e,D=me,b=V,H=E,B=Y,I=G,N=ve,A=de;return a(),f("div",{class:"post-item",onClick:s[3]||(s[3]=h=>k(e.value.id))},[n(A,{"content-indented":""},Q({avatar:o(()=>[n(_,{round:"",size:30,src:e.value.user.avatar},null,8,["src"])]),header:o(()=>[i("span",Ce,[n(j,{onClick:s[0]||(s[0]=m(()=>{},["stop"])),class:"username-link",to:{name:"user",query:{s:e.value.user.username}}},{default:o(()=>[r(p(e.value.user.nickname),1)]),_:1},8,["to"])]),i("span",$e," @"+p(e.value.user.username),1),e.value.is_top?(a(),u(y,{key:0,class:"top-tag",type:"warning",size:"small",round:""},{default:o(()=>[r(" 置顶 ")]),_:1})):c("",!0),e.value.visibility==1?(a(),u(y,{key:1,class:"top-tag",type:"error",size:"small",round:""},{default:o(()=>[r(" 私密 ")]),_:1})):c("",!0),e.value.visibility==2?(a(),u(y,{key:2,class:"top-tag",type:"info",size:"small",round:""},{default:o(()=>[r(" 好友可见 ")]),_:1})):c("",!0)]),"header-extra":o(()=>[i("div",ze,[i("span",Te,p(e.value.ip_loc?e.value.ip_loc+" · ":e.value.ip_loc)+" "+p(l(ee)(e.value.created_on)),1),n(D,{placement:"bottom-end",trigger:"hover",size:"small",options:q.value,onSelect:P},{default:o(()=>[n(M,{quaternary:"",circle:""},{icon:o(()=>[n(l(v),null,{default:o(()=>[n(l(ne))]),_:1})]),_:1})]),_:1},8,["options"])])]),footer:o(()=>[e.value.attachments.length>0?(a(),u(b,{key:0,attachments:e.value.attachments},null,8,["attachments"])):c("",!0),e.value.charge_attachments.length>0?(a(),u(b,{key:1,attachments:e.value.charge_attachments,price:e.value.attachment_price},null,8,["attachments","price"])):c("",!0),e.value.imgs.length>0?(a(),u(H,{key:2,imgs:e.value.imgs},null,8,["imgs"])):c("",!0),e.value.videos.length>0?(a(),u(B,{key:3,videos:e.value.videos},null,8,["videos"])):c("",!0),e.value.links.length>0?(a(),u(I,{key:4,links:e.value.links},null,8,["links"])):c("",!0)]),action:o(()=>[n(N,{justify:"space-between"},{default:o(()=>[i("div",{class:"opt-item hover",onClick:m(L,["stop"])},[n(l(v),{size:"18",class:"opt-item-icon"},{default:o(()=>[n(l(ae))]),_:1}),r(" "+p(e.value.upvote_count),1)],8,Pe),i("div",{class:"opt-item hover",onClick:s[2]||(s[2]=m(h=>k(e.value.id),["stop"]))},[n(l(v),{size:"18",class:"opt-item-icon"},{default:o(()=>[n(l(le))]),_:1}),r(" "+p(e.value.comment_count),1)]),i("div",{class:"opt-item hover",onClick:m(O,["stop"])},[n(l(v),{size:"18",class:"opt-item-icon"},{default:o(()=>[n(l(ie))]),_:1}),r(" "+p(e.value.collection_count),1)],8,Le)]),_:1})]),_:2},[e.value.texts.length>0?{name:"description",fn:o(()=>[(a(!0),f(U,null,W(e.value.texts,h=>(a(),f("span",{key:h.id,class:"post-text hover",onClick:s[1]||(s[1]=m(F=>S(F,e.value.id),["stop"])),innerHTML:l(R)(h.content).content},null,8,qe))),128))]),key:"0"}:void 0]),1024)])}}});export{Ae as _,Ne as a};
