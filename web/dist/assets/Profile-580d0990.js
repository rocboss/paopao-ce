import{_ as ge}from"./whisper-3ea0f829.js";import{_ as we,a as ke}from"./post-item.vue_vue_type_style_index_0_lang-f4681b38.js";import{_ as ye}from"./post-skeleton-f8a7d430.js";import{_ as be}from"./main-nav.vue_vue_type_style_index_0_lang-202c8ebb.js";import{d as Ie,H as i,b as Pe,E as Oe,r as Te,f as n,k as r,bf as u,q as f,w as _,Y as d,e as a,j as p,x as O,A as W,y as ae,F as b,u as I}from"./@vue-a481fc63.js";import{u as Fe}from"./vuex-44de225f.js";import{b as Ae}from"./vue-router-e5a2430e.js";import{e as x,J as Me,u as xe,f as ze,_ as $e}from"./index-043b20da.js";import{W as qe}from"./v3-infinite-loading-2c58ec2f.js";import{F as Ce,G as Se,a as Le,o as Ue,M as Be,f as De,g as He,J as Ne,k as Ve,H as We}from"./naive-ui-defd0b2d.js";import"./content-70c346cf.js";import"./@vicons-4ed09704.js";import"./paopao-video-player-2fe58954.js";import"./copy-to-clipboard-4ef7d3eb.js";import"./@babel-725317a4.js";import"./toggle-selection-93f4ad84.js";import"./vooks-6d99783e.js";import"./evtd-b614532e.js";import"./axios-4a70c6fc.js";import"./moment-2ab8298d.js";/* empty css               */import"./seemly-76b7b838.js";import"./vueuc-39372edb.js";import"./@css-render-7124a1a5.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";const je={class:"profile-baseinfo"},Ee={class:"avatar"},Je={class:"base-info"},Re={class:"username"},Ge={class:"userinfo"},Ye={class:"info-item"},Ke={class:"info-item"},Qe={class:"userinfo"},Xe={class:"info-item"},Ze={class:"info-item"},et={key:0,class:"skeleton-wrap"},tt={key:1},at={key:0,class:"empty-wrap"},st={key:1},lt={key:0},ot={key:1},nt={key:2},ut={key:3},it={key:4},rt={key:2},ct={key:0},_t={key:1},vt={key:2},dt={key:3},mt={key:4},ft={class:"load-more-wrap"},pt={class:"load-more-spinner"},ht=Ie({__name:"Profile",setup(gt){const o=Fe(),T=Ae(),se=Ce(),v=i(!1),P=i(!1),l=i([]),z=i([]),$=i([]),q=i([]),C=i([]),S=i([]),m=i("post"),j=i(+T.query.p||1),E=i(1),J=i(1),R=i(1),G=i(1),s=i(+T.query.p||1),g=i(20),c=i(0),Y=i(0),K=i(0),Q=i(0),X=i(0),Z=i(0),L=i(!1),ee=i({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1}),w=e=>{ee.value=e,L.value=!0},le=()=>{L.value=!1},k=e=>{se.success({title:"提示",content:"确定"+(e.user.is_following?"取消关注":"关注")+"该用户吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{e.user.is_following?xe({user_id:e.user.id}).then(h=>{window.$message.success("操作成功"),e.user.is_following=!1}).catch(h=>{}):ze({user_id:e.user.id}).then(h=>{window.$message.success("关注成功"),e.user.is_following=!0}).catch(h=>{})}})},U=()=>{switch(m.value){case"post":B();break;case"comment":D();break;case"highlight":H();break;case"media":N();break;case"star":V();break}},B=()=>{v.value=!0,x({username:o.state.userInfo.username,style:"post",page:s.value,page_size:g.value}).then(e=>{v.value=!1,e.list.length===0&&(P.value=!0),s.value>1?l.value=l.value.concat(e.list):(l.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/g.value),z.value=l.value,Y.value=c.value}).catch(e=>{l.value=[],s.value>1&&s.value--,v.value=!1})},D=()=>{v.value=!0,x({username:o.state.userInfo.username,style:"comment",page:s.value,page_size:g.value}).then(e=>{v.value=!1,e.list.length===0&&(P.value=!0),s.value>1?l.value=l.value.concat(e.list):(l.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/g.value),$.value=l.value,K.value=c.value}).catch(e=>{l.value=[],s.value>1&&s.value--,v.value=!1})},H=()=>{v.value=!0,x({username:o.state.userInfo.username,style:"highlight",page:s.value,page_size:g.value}).then(e=>{v.value=!1,e.list.length===0&&(P.value=!0),s.value>1?l.value=l.value.concat(e.list):(l.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/g.value),q.value=l.value,Q.value=c.value}).catch(e=>{l.value=[],s.value>1&&s.value--,v.value=!1})},N=()=>{v.value=!0,x({username:o.state.userInfo.username,style:"media",page:s.value,page_size:g.value}).then(e=>{v.value=!1,e.list.length===0&&(P.value=!0),s.value>1?l.value=l.value.concat(e.list):(l.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/g.value),C.value=l.value,X.value=c.value}).catch(e=>{l.value=[],s.value>1&&s.value--,v.value=!1})},V=()=>{v.value=!0,x({username:o.state.userInfo.username,style:"star",page:s.value,page_size:g.value}).then(e=>{v.value=!1,e.list.length===0&&(P.value=!0),s.value>1?l.value=l.value.concat(e.list):(l.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/g.value),S.value=l.value,Z.value=c.value}).catch(e=>{l.value=[],s.value>1&&s.value--,v.value=!1})},oe=e=>{switch(m.value=e,m.value){case"post":l.value=z.value,s.value=j.value,c.value=Y.value,B();break;case"comment":l.value=$.value,s.value=E.value,c.value=K.value,D();break;case"highlight":l.value=q.value,s.value=J.value,c.value=Q.value,H();break;case"media":l.value=C.value,s.value=R.value,c.value=X.value,N();break;case"star":l.value=S.value,s.value=G.value,c.value=Z.value,V();break}},ne=()=>{switch(m.value){case"post":j.value=s.value,B();break;case"comment":E.value=s.value,D();break;case"highlight":J.value=s.value,H();break;case"media":R.value=s.value,N();break;case"star":G.value=s.value,V();break}},ue=()=>{s.value<c.value||c.value==0?(P.value=!1,s.value++,ne()):P.value=!0};return Pe(()=>{U()}),Oe(()=>({path:T.path,query:T.query,refresh:o.state.refresh}),(e,h)=>{if(e.refresh!==h.refresh){s.value=+T.query.p||1,setTimeout(()=>{U()},0);return}h.path!=="/post"&&e.path==="/profile"&&(s.value=+T.query.p||1,setTimeout(()=>{U()},0))}),(e,h)=>{const ie=be,re=Ue,ce=Be,te=Te("router-link"),F=De,_e=He,ve=ye,de=Ne,A=we,y=We,M=ke,me=ge,fe=Se,pe=Ve,he=Le;return a(),n("div",null,[r(ie,{title:"主页"}),u(o).state.userInfo.id>0?(a(),f(fe,{key:0,class:"main-content-wrap profile-wrap",bordered:""},{default:_(()=>[p("div",je,[p("div",Ee,[r(re,{size:72,src:u(o).state.userInfo.avatar},null,8,["src"])]),p("div",Je,[p("div",Re,[p("strong",null,O(u(o).state.userInfo.nickname),1),p("span",null," @"+O(u(o).state.userInfo.username),1),u(o).state.userInfo.is_admin?(a(),f(ce,{key:0,class:"top-tag",type:"error",size:"small",round:""},{default:_(()=>[W(" 管理员 ")]),_:1})):d("",!0)]),p("div",Ge,[p("span",Ye,"UID. "+O(u(o).state.userInfo.id),1),p("span",Ke,O(u(Me)(u(o).state.userInfo.created_on))+" 加入",1)]),p("div",Qe,[p("span",Xe,[r(te,{onClick:h[0]||(h[0]=ae(()=>{},["stop"])),class:"following-link",to:{name:"following",query:{s:u(o).state.userInfo.username,n:u(o).state.userInfo.nickname,t:"follows"}}},{default:_(()=>[W(" 关注  "+O(u(o).state.userInfo.follows),1)]),_:1},8,["to"])]),p("span",Ze,[r(te,{onClick:h[1]||(h[1]=ae(()=>{},["stop"])),class:"following-link",to:{name:"following",query:{s:u(o).state.userInfo.username,n:u(o).state.userInfo.nickname,t:"followings"}}},{default:_(()=>[W(" 粉丝  "+O(u(o).state.userInfo.followings),1)]),_:1},8,["to"])])])])]),r(_e,{class:"profile-tabs-wrap",type:"line",animated:"","onUpdate:value":oe},{default:_(()=>[r(F,{name:"post",tab:"泡泡"}),r(F,{name:"comment",tab:"评论"}),r(F,{name:"highlight",tab:"亮点"}),r(F,{name:"media",tab:"图文"}),r(F,{name:"star",tab:"喜欢"})]),_:1}),v.value&&l.value.length===0?(a(),n("div",et,[r(ve,{num:g.value},null,8,["num"])])):(a(),n("div",tt,[l.value.length===0?(a(),n("div",at,[r(de,{size:"large",description:"暂无数据"})])):d("",!0),u(o).state.desktopModelShow?(a(),n("div",st,[m.value==="post"?(a(),n("div",lt,[(a(!0),n(b,null,I(z.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(A,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0),m.value==="comment"?(a(),n("div",ot,[(a(!0),n(b,null,I($.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(A,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0),m.value==="highlight"?(a(),n("div",nt,[(a(!0),n(b,null,I(q.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(A,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0),m.value==="media"?(a(),n("div",ut,[(a(!0),n(b,null,I(C.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(A,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0),m.value==="star"?(a(),n("div",it,[(a(!0),n(b,null,I(S.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(A,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0)])):(a(),n("div",rt,[m.value==="post"?(a(),n("div",ct,[(a(!0),n(b,null,I(z.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(M,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0),m.value==="comment"?(a(),n("div",_t,[(a(!0),n(b,null,I($.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(M,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0),m.value==="highlight"?(a(),n("div",vt,[(a(!0),n(b,null,I(q.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(M,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0),m.value==="media"?(a(),n("div",dt,[(a(!0),n(b,null,I(C.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(M,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0),m.value==="star"?(a(),n("div",mt,[(a(!0),n(b,null,I(S.value,t=>(a(),f(y,{key:t.id},{default:_(()=>[r(M,{post:t,isOwner:u(o).state.userInfo.id==t.user_id,addFollowAction:!0,onSendWhisper:w,onHandleFollowAction:k},null,8,["post","isOwner"])]),_:2},1024))),128))])):d("",!0)]))])),r(me,{show:L.value,user:ee.value,onSuccess:le},null,8,["show","user"])]),_:1})):d("",!0),c.value>0?(a(),f(he,{key:1,justify:"center"},{default:_(()=>[r(u(qe),{class:"load-more",slots:{complete:"没有更多泡泡了",error:"加载出错"},onInfinite:h[2]||(h[2]=t=>ue())},{spinner:_(()=>[p("div",ft,[P.value?d("",!0):(a(),f(pe,{key:0,size:14})),p("span",pt,O(P.value?"没有更多泡泡了":"加载更多"),1)])]),_:1})]),_:1})):d("",!0)])}}});const Qt=$e(ht,[["__scopeId","data-v-4727fe2e"]]);export{Qt as default};