import{d as ae,r,l as we,f as _,j as o,k as m,q as t,t as z,Y as u,w as n,G as O,z as A,bq as a,A as Y,c as ce,ar as Je,F as ve,x as fe,b as Re,h as Te,C as Ke}from"./@vue-9sINKCPW.js";import{u as re}from"./vuex--ttreJMD.js";import{i as he,t as We,j as Qe,k as Xe,_ as pe,l as Ze,m as et,n as tt,o as st,p as ot,g as nt,q as at,r as lt,s as it,v as ut,w as ct,x as rt,y as pt,z as _t,u as dt,f as mt,A as vt,B as ft,C as Pe,D as ke}from"./index-DxHQoSDp.js";import{Y as se,V as ee}from"./IEnum-B3rDUvtK.js";import{K as Se,B as _e,i as K,I as ht,L as gt,e as yt,M as Oe,O as Le,m as be,t as kt,v as wt,w as bt,x as $t,y as Ct,z as xt,F as Tt,a as Ae,P as Pt,Q as Ut,_ as zt,G as It,H as Rt,j as St,J as Ot,f as Lt,g as At}from"./naive-ui-BJojRuLw.js";import{T as De,e as ge,f as Me,g as ye,h as qe,i as Dt,j as Mt,I as qt,k as Nt,l as Et,m as Bt,n as Ft,o as Ht,p as jt,q as Yt,r as Vt,s as Gt,t as ue,u as Jt,v as Kt,w as Wt,x as Ue,F as ze,y as de,E as me,z as Ie}from"./@vicons-PCg97L0F.js";import{p as $e,_ as Ne,a as Qt,b as Xt,c as Zt}from"./content-x8CCof-x.js";import{u as Ee,b as es}from"./vue-router-yrkFRUM9.js";import{_ as ts}from"./post-skeleton-B6KFVL2X.js";import{l as ss}from"./lodash-CGvuAYz8.js";import{_ as os}from"./whisper-CD0RDnNY.js";import{a as ns}from"./copy-to-clipboard-DgsYVcxl.js";import{_ as as}from"./main-nav.vue_vue_type_style_index_0_lang-D7-FAbTY.js";import{Y as ls}from"./v3-infinite-loading-D3303HHP.js";import"./axios-t--hEgTQ.js";import"./moment-P60zs0je.js";/* empty css               */import"./seemly-D-teBmey.js";import"./vueuc-DrMWnH2h.js";import"./evtd-CI_DDEu_.js";import"./@css-render-DN2R7sM6.js";import"./vooks-CfQnrjIt.js";import"./vdirs-DRH9Xvnd.js";import"./@juggle-C8OzoCMD.js";import"./css-render-BDrvWz3H.js";import"./@emotion-WldOFDRm.js";import"./lodash-es-TJvrUncL.js";import"./treemate-HRdUPn5m.js";import"./async-validator-9PlIezaS.js";import"./date-fns-Db9XENWt.js";import"./paopao-video-player-FrzfkELx.js";import"./toggle-selection-DGa8lynz.js";const is={class:"reply-item"},us={class:"header-wrap"},cs={class:"username"},rs={class:"reply-name"},ps={class:"timestamp"},_s={class:"base-wrap"},ds={class:"content"},ms={class:"reply-switch"},vs={class:"time-item"},fs={class:"actions"},hs={class:"upvote-count"},gs={class:"upvote-count"},ys={key:2,class:"action-item"},ks=ae({__name:"reply-item",props:{tweetId:{},reply:{}},emits:["focusReply","reload"],setup(J,{emit:j}){const c=J,d=re(),k=j,x=r(c.reply.is_thumbs_up==se.YES),L=r(c.reply.is_thumbs_down==se.YES),U=r(c.reply.thumbs_up_count),f=()=>{We({tweet_id:c.tweetId,comment_id:c.reply.comment_id,reply_id:c.reply.id}).then(g=>{x.value=!x.value,x.value?(U.value++,L.value=!1):U.value--}).catch(g=>{console.log(g)})},l=()=>{Qe({tweet_id:c.tweetId,comment_id:c.reply.comment_id,reply_id:c.reply.id}).then(g=>{L.value=!L.value,L.value&&x.value&&(U.value--,x.value=!1)}).catch(g=>{console.log(g)})},q=()=>{k("focusReply",c.reply)},D=()=>{Xe({id:c.reply.id}).then(g=>{window.$message.success("删除成功"),setTimeout(()=>{k("reload")},50)}).catch(g=>{console.log(g)})};return(g,$)=>{const M=we("router-link"),P=K,C=_e,w=Se,b=ht;return o(),_("div",is,[m("div",us,[m("div",cs,[t(M,{class:"user-link",to:{name:"user",query:{s:c.reply.user.username}}},{default:n(()=>[O(A(c.reply.user.username),1)]),_:1},8,["to"]),m("span",rs,A(c.reply.at_user_id>0?"回复":":"),1),c.reply.at_user_id>0?(o(),z(M,{key:0,class:"user-link",to:{name:"user",query:{s:c.reply.at_user.username}}},{default:n(()=>[O(A(c.reply.at_user.username),1)]),_:1},8,["to"])):u("",!0)]),m("div",ps,[O(A(c.reply.ip_loc)+" ",1),a(d).state.userInfo.is_admin||a(d).state.userInfo.id===c.reply.user.id?(o(),z(w,{key:0,"negative-text":"取消","positive-text":"确认",onPositiveClick:D},{trigger:n(()=>[t(C,{quaternary:"",circle:"",size:"tiny",class:"del-btn"},{icon:n(()=>[t(P,null,{default:n(()=>[t(a(De))]),_:1})]),_:1})]),default:n(()=>[$[1]||($[1]=O(" 是否删除这条回复？ "))]),_:1})):u("",!0)])]),m("div",_s,[m("div",ds,[t(b,{"expand-trigger":"click","line-clamp":"5",tooltip:!1},{default:n(()=>[O(A(c.reply.content),1)]),_:1})]),m("div",ms,[m("span",vs,A(a(he)(c.reply.created_on)),1),m("div",fs,[a(d).state.userLogined?u("",!0):(o(),_("div",{key:0,class:"action-item",onClick:$[0]||($[0]=Y(()=>{},["stop"]))},[t(P,{size:"medium"},{default:n(()=>[t(a(ge))]),_:1}),m("span",hs,A(U.value),1)])),a(d).state.userLogined?(o(),_("div",{key:1,class:"action-item hover",onClick:Y(f,["stop"])},[t(P,{size:"medium"},{default:n(()=>[x.value?u("",!0):(o(),z(a(ge),{key:0})),x.value?(o(),z(a(Me),{key:1,class:"show"})):u("",!0)]),_:1}),m("span",gs,A(U.value>0?U.value:"赞"),1)])):u("",!0),a(d).state.userLogined?u("",!0):(o(),_("div",ys,[t(P,{size:"medium"},{default:n(()=>[t(a(ye))]),_:1})])),a(d).state.userLogined?(o(),_("div",{key:3,class:"action-item hover",onClick:Y(l,["stop"])},[t(P,{size:"medium"},{default:n(()=>[L.value?u("",!0):(o(),z(a(ye),{key:0})),L.value?(o(),z(a(qe),{key:1,class:"show"})):u("",!0)]),_:1})])):u("",!0),a(d).state.userLogined?(o(),_("span",{key:4,class:"show opacity-item reply-btn",onClick:q}," 回复 ")):u("",!0)])])])])}}}),ws=pe(ks,[["__scopeId","data-v-eccdbbd8"]]),bs={class:"reply-compose-wrap"},$s={class:"reply-switch"},Cs={class:"time-item"},xs={class:"actions"},Ts={key:0,class:"action-item"},Ps={class:"upvote-count"},Us={class:"upvote-count"},zs={key:2,class:"action-item"},Is={key:0,class:"reply-input-wrap"},Rs=ae({__name:"compose-reply",props:{comment:{},atUserid:{default:0},atUsername:{default:""}},emits:["reload","reset"],setup(J,{expose:j,emit:c}){const d=J,k=re(),x=c,L=r(),U=r(!1),f=r(""),l=r(!1),q=300,D=r(d.comment.is_thumbs_up==se.YES),g=r(d.comment.is_thumbs_down==se.YES),$=r(d.comment.thumbs_up_count),M=()=>{Ze({tweet_id:d.comment.post_id,comment_id:d.comment.id}).then(b=>{D.value=!D.value,D.value?($.value++,g.value=!1):$.value--}).catch(b=>{console.log(b)})},P=()=>{et({tweet_id:d.comment.post_id,comment_id:d.comment.id}).then(b=>{g.value=!g.value,g.value&&D.value&&($.value--,D.value=!1)}).catch(b=>{console.log(b)})},C=b=>{U.value=b,b?setTimeout(()=>{var I;(I=L.value)==null||I.focus()},10):(l.value=!1,f.value="",x("reset"))},w=()=>{l.value=!0,tt({comment_id:d.comment.id,at_user_id:d.atUserid,content:f.value}).then(b=>{C(!1),window.$message.success("评论成功"),x("reload")}).catch(b=>{l.value=!1})};return j({switchReply:C}),(b,I)=>{const N=K,s=yt,v=_e,E=gt;return o(),_("div",bs,[m("div",$s,[m("span",Cs,A(a(he)(b.comment.created_on)),1),m("div",xs,[a(k).state.userLogined?u("",!0):(o(),_("div",Ts,[t(N,{size:"medium"},{default:n(()=>[t(a(ge))]),_:1}),m("span",Ps,A($.value),1)])),a(k).state.userLogined?(o(),_("div",{key:1,class:"action-item hover",onClick:Y(M,["stop"])},[t(N,{size:"medium"},{default:n(()=>[D.value?u("",!0):(o(),z(a(ge),{key:0})),D.value?(o(),z(a(Me),{key:1,class:"show"})):u("",!0)]),_:1}),m("span",Us,A($.value>0?$.value:"赞"),1)])):u("",!0),a(k).state.userLogined?u("",!0):(o(),_("div",zs,[t(N,{size:"medium"},{default:n(()=>[t(a(ye))]),_:1})])),a(k).state.userLogined?(o(),_("div",{key:3,class:"action-item hover",onClick:Y(P,["stop"])},[t(N,{size:"medium"},{default:n(()=>[g.value?u("",!0):(o(),z(a(ye),{key:0})),g.value?(o(),z(a(qe),{key:1,class:"show"})):u("",!0)]),_:1})])):u("",!0),a(k).state.userLogined&&!U.value?(o(),_("span",{key:4,class:"show reply-btn",onClick:I[0]||(I[0]=V=>C(!0))}," 回复 ")):u("",!0),a(k).state.userLogined&&U.value?(o(),_("span",{key:5,class:"hide reply-btn",onClick:I[1]||(I[1]=V=>C(!1))}," 取消 ")):u("",!0)])]),U.value?(o(),_("div",Is,[t(E,null,{default:n(()=>[t(s,{ref_key:"inputInstRef",ref:L,size:"small",placeholder:d.atUsername?"@"+d.atUsername:"请输入回复内容..",maxlength:a(q),value:f.value,"onUpdate:value":I[2]||(I[2]=V=>f.value=V),"show-count":"",clearable:""},null,8,["placeholder","maxlength","value"]),t(v,{type:"primary",size:"small",ghost:"",loading:l.value,onClick:w},{default:n(()=>I[3]||(I[3]=[O(" 回复 ")])),_:1},8,["loading"])]),_:1})])):u("",!0)])}}}),Ss=pe(Rs,[["__scopeId","data-v-f9af7a93"]]),Os={class:"comment-item"},Ls={class:"nickname-wrap"},As={class:"username-wrap"},Ds={class:"opt-wrap"},Ms={class:"timestamp"},qs=["innerHTML"],Ns={class:"reply-wrap"},Es=ae({__name:"comment-item",props:{comment:{},postUserId:{}},emits:["reload"],setup(J,{emit:j}){const c=re(),d=Ee(),k=r(0),x=r(""),L=r(),U=j,f=J,l=ce(()=>{let C=Object.assign({texts:[],imgs:[]},f.comment);return C.contents.map(w=>{(+w.type==1||+w.type==2)&&C.texts.push(w),+w.type==3&&C.imgs.push(w)}),C}),q=(C,w)=>{let b=C.target;if(b.dataset.detail){const I=b.dataset.detail.split(":");I.length===2&&(c.commit("refresh"),I[0]==="tag"?window.$message.warning("评论内的无效话题"):d.push({name:"user",query:{s:I[1]}}))}},D=C=>{var w,b;k.value=C.user_id,x.value=((w=C.user)==null?void 0:w.username)||"",(b=L.value)==null||b.switchReply(!0)},g=()=>{U("reload")},$=()=>{k.value=0,x.value=""},M=()=>{st({id:l.value.id}).then(C=>{window.$message.success("删除成功"),setTimeout(()=>{g()},50)}).catch(C=>{})},P=()=>{ot({id:l.value.id}).then(C=>{l.value.is_essence=C.highlight_status,window.$message.success("操作成功"),setTimeout(()=>{g()},50)}).catch(C=>{})};return(C,w)=>{const b=be,I=we("router-link"),N=Le,s=K,v=_e,E=Se,V=Ne,G=Ss,W=ws,R=Oe;return o(),_("div",Os,[t(R,{"content-indented":""},Je({avatar:n(()=>[t(b,{round:"",size:30,src:l.value.user.avatar},null,8,["src"])]),header:n(()=>[m("span",Ls,[t(I,{onClick:w[0]||(w[0]=Y(()=>{},["stop"])),class:"username-link",to:{name:"user",query:{s:l.value.user.username}}},{default:n(()=>[O(A(l.value.user.nickname),1)]),_:1},8,["to"])]),m("span",As," @"+A(l.value.user.username),1),l.value.is_essence==a(se).YES?(o(),z(N,{key:0,class:"top-tag",type:"warning",size:"small",round:""},{default:n(()=>w[2]||(w[2]=[O(" 精选 ")])),_:1})):u("",!0)]),"header-extra":n(()=>[m("div",Ds,[m("span",Ms,A(l.value.ip_loc),1),a(c).state.userInfo.id===C.postUserId?(o(),z(E,{key:0,"negative-text":"取消","positive-text":"确认",onPositiveClick:P},{trigger:n(()=>[t(v,{quaternary:"",circle:"",size:"tiny",class:"action-btn"},{icon:n(()=>[l.value.is_essence==a(se).NO?(o(),z(s,{key:0},{default:n(()=>[t(a(Dt))]),_:1})):(o(),z(s,{key:1},{default:n(()=>[t(a(Mt))]),_:1}))]),_:1})]),default:n(()=>[O(" "+A(l.value.is_essence==a(se).NO?"是否精选这条评论":"是否取消精选"),1)]),_:1})):u("",!0),a(c).state.userInfo.is_admin||a(c).state.userInfo.id===l.value.user.id?(o(),z(E,{key:1,"negative-text":"取消","positive-text":"确认",onPositiveClick:M},{trigger:n(()=>[t(v,{quaternary:"",circle:"",size:"tiny",class:"action-btn"},{icon:n(()=>[t(s,null,{default:n(()=>[t(a(De))]),_:1})]),_:1})]),default:n(()=>[w[3]||(w[3]=O(" 是否删除这条评论？ "))]),_:1})):u("",!0)])]),footer:n(()=>[l.value.imgs.length>0?(o(),z(V,{key:0,imgs:l.value.imgs},null,8,["imgs"])):u("",!0),t(G,{ref_key:"replyComposeRef",ref:L,comment:l.value,"at-userid":k.value,"at-username":x.value,onReload:g,onReset:$},null,8,["comment","at-userid","at-username"]),m("div",Ns,[(o(!0),_(ve,null,fe(l.value.replies,y=>(o(),z(W,{key:y.id,reply:y,"tweet-id":l.value.post_id,onFocusReply:D,onReload:g},null,8,["reply","tweet-id"]))),128))])]),_:2},[l.value.texts.length>0?{name:"description",fn:n(()=>[(o(!0),_(ve,null,fe(l.value.texts,y=>(o(),_("span",{key:y.id,class:"comment-text",onClick:w[1]||(w[1]=Y(Q=>q(Q,l.value.id),["stop"])),innerHTML:a($e)(y.content).content},null,8,qs))),128))]),key:"0"}:void 0]),1024)])}}}),Bs=pe(Es,[["__scopeId","data-v-e1f04c6b"]]),Fs={key:0,class:"compose-wrap"},Hs={class:"compose-line"},js={class:"compose-user"},Ys={class:"compose-line compose-options"},Vs={class:"attachment"},Gs={class:"submit-wrap"},Js={class:"attachment-list-wrap"},Ks={key:1,class:"compose-wrap"},Ws={key:0,class:"login-only-wrap"},Qs={key:1,class:"login-wrap"},Xs=ae({__name:"compose-comment",props:{lock:{default:0},postId:{default:0}},emits:["post-success"],setup(J,{emit:j}){const c=j,d=J,k=re(),x=r([]),L=r(!1),U=r(!1),f=r(!1),l=r(""),q=r(),D=r("public/image"),g=r([]),$=r([]),M=r("true".toLowerCase()==="true"),P=300,C="/v1/attachment",w=ce(()=>"Bearer "+localStorage.getItem("PAOPAO_TOKEN")),b=ss.debounce(h=>{nt({k:h}).then(p=>{let T=[];p.suggest.map(S=>{T.push({label:S,value:S})}),x.value=T,U.value=!1}).catch(p=>{U.value=!1})},200),I=(h,p)=>{U.value||(U.value=!0,p==="@"&&b(h))},N=h=>{h.length>P?l.value=h.substring(0,P):l.value=h},s=h=>{D.value=h},v=h=>{for(let B=0;B<h.length;B++){var p=h[B].name,T=p.split(".").slice(0,-1).join("."),S=p.split(".").pop();T.length>30&&(h[B].name=T.substring(0,18)+"..."+T.substring(T.length-9)+"."+S)}g.value=h},E=async h=>{var p,T;return D.value==="public/image"&&!["image/png","image/jpg","image/jpeg","image/gif"].includes((p=h.file.file)==null?void 0:p.type)?(window.$message.warning("图片仅允许 png/jpg/gif 格式"),!1):D.value==="image"&&((T=h.file.file)==null?void 0:T.size)>10485760?(window.$message.warning("图片大小不能超过10MB"),!1):!0},V=({file:h,event:p})=>{var T;try{let S=JSON.parse((T=p.target)==null?void 0:T.response);S.code===0&&D.value==="public/image"&&$.value.push({id:h.id,content:S.data.content})}catch{window.$message.error("上传失败")}},G=({file:h,event:p})=>{var T;try{let S=JSON.parse((T=p.target)==null?void 0:T.response);if(S.code!==0){let B=S.msg||"上传失败";S.details&&S.details.length>0&&S.details.map(e=>{B+=":"+e}),window.$message.error(B)}}catch{window.$message.error("上传失败")}},W=({file:h})=>{let p=$.value.findIndex(T=>T.id===h.id);p>-1&&$.value.splice(p,1)},R=()=>{L.value=!0},y=()=>{var h;L.value=!1,(h=q.value)==null||h.clear(),g.value=[],l.value="",$.value=[]},Q=()=>{if(l.value.trim().length===0){window.$message.warning("请输入内容哦");return}let{users:h}=$e(l.value);const p=[];let T=100;p.push({content:l.value,type:2,sort:T}),$.value.map(S=>{T++,p.push({content:S.content,type:3,sort:T})}),f.value=!0,at({contents:p,post_id:d.postId,users:Array.from(new Set(h))}).then(S=>{window.$message.success("发布成功"),f.value=!1,c("post-success"),y()}).catch(S=>{f.value=!1})},X=h=>{k.commit("triggerAuth",!0),k.commit("triggerAuthKey",h)};return(h,p)=>{const T=be,S=kt,B=K,e=_e,i=bt,F=Ct,le=$t,Z=xt,oe=wt;return o(),_("div",null,[a(k).state.userInfo.id>0?(o(),_("div",Fs,[m("div",Hs,[m("div",js,[t(T,{round:"",size:30,src:a(k).state.userInfo.avatar},null,8,["src"])]),t(S,{type:"textarea",size:"large",autosize:"",bordered:!1,options:x.value,prefix:["@"],loading:U.value,value:l.value,disabled:d.lock===1,"onUpdate:value":N,onSearch:I,onFocus:R,placeholder:d.lock===1?"泡泡已被锁定，回复功能已关闭":"快来评论两句吧..."},null,8,["options","loading","value","disabled","placeholder"])]),L.value?(o(),z(oe,{key:0,ref_key:"uploadRef",ref:q,abstract:"","list-type":"image",multiple:!0,max:9,action:C,headers:{Authorization:w.value},data:{type:D.value},"file-list":g.value,onBeforeUpload:E,onFinish:V,onError:G,onRemove:W,"onUpdate:fileList":v},{default:n(()=>[m("div",Ys,[m("div",Vs,[t(i,{abstract:""},{default:n(({handleClick:ne})=>[t(e,{disabled:g.value.length>0&&D.value==="public/video"||g.value.length===9,onClick:()=>{s("public/image"),ne()},quaternary:"",circle:"",type:"primary"},{icon:n(()=>[t(B,{size:"20",color:"var(--primary-color)"},{default:n(()=>[t(a(qt))]),_:1})]),_:2},1032,["disabled","onClick"])]),_:1}),t(le,{trigger:"hover",placement:"bottom"},{trigger:n(()=>[t(F,{class:"text-statistic",type:"circle","show-indicator":!1,status:"success","stroke-width":10,percentage:l.value.length/a(P)*100},null,8,["percentage"])]),default:n(()=>[O(" "+A(l.value.length)+" / "+A(a(P)),1)]),_:1})]),m("div",Gs,[t(e,{quaternary:"",round:"",type:"tertiary",class:"cancel-btn",size:"small",onClick:y},{default:n(()=>p[3]||(p[3]=[O(" 取消 ")])),_:1}),t(e,{loading:f.value,onClick:Q,type:"primary",secondary:"",size:"small",round:""},{default:n(()=>p[4]||(p[4]=[O(" 发布 ")])),_:1},8,["loading"])])]),m("div",Js,[t(Z)])]),_:1},8,["headers","data","file-list"])):u("",!0)])):(o(),_("div",Ks,[p[8]||(p[8]=m("div",{class:"login-wrap"},[m("span",{class:"login-banner"}," 登录后，精彩更多")],-1)),M.value?u("",!0):(o(),_("div",Ws,[t(e,{strong:"",secondary:"",round:"",type:"primary",onClick:p[0]||(p[0]=ne=>X("signin"))},{default:n(()=>p[5]||(p[5]=[O(" 登录 ")])),_:1})])),M.value?(o(),_("div",Qs,[t(e,{strong:"",secondary:"",round:"",type:"primary",onClick:p[1]||(p[1]=ne=>X("signin"))},{default:n(()=>p[6]||(p[6]=[O(" 登录 ")])),_:1}),t(e,{strong:"",secondary:"",round:"",type:"info",onClick:p[2]||(p[2]=ne=>X("signup"))},{default:n(()=>p[7]||(p[7]=[O(" 注册 ")])),_:1})])):u("",!0)]))])}}}),Zs=pe(Xs,[["__scopeId","data-v-d9073453"]]),eo={class:"username-wrap"},to={class:"options"},so={key:0},oo=["innerHTML"],no={class:"timestamp"},ao={key:0},lo={key:1},io={class:"opts-wrap"},uo={class:"opt-item"},co=ae({__name:"post-detail",props:{post:{}},emits:["reload"],setup(J,{emit:j}){const c="true".toLowerCase()==="true",d=re(),k=Ee(),x=Tt(),L=r(!1),U=r(!1),f=J,l=r(!1),q=r(!1),D=r(!1),g=r(!1),$=r(!1),M=r(!1),P=r(ee.PUBLIC),C=r(!1),w=r({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1}),b=e=>{w.value=e,C.value=!0},I=()=>{C.value=!1},N=j,s=ce({get:()=>{let e=Object.assign({texts:[],imgs:[],videos:[],links:[],attachments:[],charge_attachments:[]},f.post);return e.contents.map(i=>{(+i.type==1||+i.type==2)&&e.texts.push(i),+i.type==3&&e.imgs.push(i),+i.type==4&&e.videos.push(i),+i.type==6&&e.links.push(i),+i.type==7&&e.attachments.push(i),+i.type==8&&e.charge_attachments.push(i)}),e},set:e=>{f.post.upvote_count=e.upvote_count,f.post.comment_count=e.comment_count,f.post.collection_count=e.collection_count,f.post.is_essence=e.is_essence}}),v=e=>()=>Te(K,null,{default:()=>Te(e)}),E=ce(()=>{var F;let e=[];if(!d.state.userInfo.is_admin&&d.state.userInfo.id!=f.post.user.id)return e.push({label:"私信 @"+f.post.user.username,key:"whisper",icon:v(Vt)}),f.post.user.is_following?e.push({label:"取消关注 @"+f.post.user.username,key:"unfollow",icon:v(Gt)}):e.push({label:"关注 @"+f.post.user.username,key:"follow",icon:v(ue)}),e;e.push({label:"删除",key:"delete",icon:v(Jt)}),s.value.is_lock===0?e.push({label:"锁定",key:"lock",icon:v(Kt)}):e.push({label:"解锁",key:"unlock",icon:v(Wt)}),d.state.userInfo.is_admin&&(s.value.is_top===0?e.push({label:"置顶",key:"stick",icon:v(Ue)}):e.push({label:"取消置顶",key:"unstick",icon:v(Ue)})),s.value.is_essence===0?e.push({label:"设为亮点",key:"highlight",icon:v(ze)}):e.push({label:"取消亮点",key:"unhighlight",icon:v(ze)});let i;return s.value.visibility===ee.PUBLIC?i={label:"公开",key:"vpublic",icon:v(me),children:[{label:"私密",key:"vprivate",icon:v(de)},{label:"关注可见",key:"vfollowing",icon:v(ue)}]}:s.value.visibility===ee.PRIVATE?i={label:"私密",key:"vprivate",icon:v(de),children:[{label:"公开",key:"vpublic",icon:v(me)},{label:"关注可见",key:"vfollowing",icon:v(ue)}]}:c&&s.value.visibility===ee.FRIEND?i={label:"好友可见",key:"vfriend",icon:v(Ie),children:[{label:"公开",key:"vpublic",icon:v(me)},{label:"私密",key:"vprivate",icon:v(de)},{label:"关注可见",key:"vfollowing",icon:v(ue)}]}:i={label:"关注可见",key:"vfollowing",icon:v(ue),children:[{label:"公开",key:"vpublic",icon:v(me)},{label:"私密",key:"vprivate",icon:v(de)}]},c&&s.value.visibility!==ee.FRIEND&&((F=i.children)==null||F.push({label:"好友可见",key:"vfriend",icon:v(Ie)})),e.push(i),e}),V=e=>{x.success({title:"提示",content:"确定"+(e.user.is_following?"取消关注 @":"关注 @")+f.post.user.username+" 吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{e.user.is_following?dt({user_id:e.user.id}).then(i=>{window.$message.success("操作成功"),e.user.is_following=!1}).catch(i=>{}):mt({user_id:e.user.id}).then(i=>{window.$message.success("操作成功"),e.user.is_following=!0}).catch(i=>{})}})},G=e=>{k.push({name:"post",query:{id:e}})},W=(e,i)=>{if(e.target.dataset.detail){const F=e.target.dataset.detail.split(":");if(F.length===2){d.commit("refresh"),F[0]==="tag"?k.push({name:"home",query:{q:F[1],t:"tag"}}):k.push({name:"user",query:{s:F[1]}});return}}G(i)},R=e=>{switch(e){case"whisper":b(f.post.user);break;case"follow":case"unfollow":V(f.post);break;case"delete":l.value=!0;break;case"lock":case"unlock":q.value=!0;break;case"stick":case"unstick":D.value=!0;break;case"highlight":case"unhighlight":g.value=!0;break;case"vpublic":P.value=0,$.value=!0;break;case"vprivate":P.value=1,$.value=!0;break;case"vfriend":P.value=2,$.value=!0;break;case"vfollowing":P.value=3,$.value=!0;break}},y=()=>{ut({id:s.value.id}).then(e=>{window.$message.success("删除成功"),k.replace("/"),setTimeout(()=>{d.commit("refresh")},50)}).catch(e=>{M.value=!1})},Q=()=>{ct({id:s.value.id}).then(e=>{N("reload",s.value.id),e.lock_status===1?window.$message.success("锁定成功"):window.$message.success("解锁成功")}).catch(e=>{M.value=!1})},X=()=>{rt({id:s.value.id}).then(e=>{N("reload",s.value.id),e.top_status===1?window.$message.success("置顶成功"):window.$message.success("取消置顶成功")}).catch(e=>{M.value=!1})},h=()=>{pt({id:s.value.id}).then(e=>{s.value={...s.value,is_essence:e.highlight_status},e.highlight_status===1?window.$message.success("设为亮点成功"):window.$message.success("取消亮点成功")}).catch(e=>{M.value=!1})},p=()=>{_t({id:s.value.id,visibility:P.value}).then(e=>{N("reload",s.value.id),window.$message.success("修改可见性成功")}).catch(e=>{M.value=!1})},T=()=>{vt({id:s.value.id}).then(e=>{L.value=e.status,e.status?s.value={...s.value,upvote_count:s.value.upvote_count+1}:s.value={...s.value,upvote_count:s.value.upvote_count-1}}).catch(e=>{console.log(e)})},S=()=>{ft({id:s.value.id}).then(e=>{U.value=e.status,e.status?s.value={...s.value,collection_count:s.value.collection_count+1}:s.value={...s.value,collection_count:s.value.collection_count-1}}).catch(e=>{console.log(e)})},B=()=>{ns(`${window.location.origin}/#/post?id=${s.value.id}&share=copy_link&t=${new Date().getTime()}`),window.$message.success("链接已复制到剪贴板")};return Re(()=>{d.state.userInfo.id>0&&(lt({id:s.value.id}).then(e=>{L.value=e.status}).catch(e=>{console.log(e)}),it({id:s.value.id}).then(e=>{U.value=e.status}).catch(e=>{console.log(e)}))}),(e,i)=>{const F=be,le=we("router-link"),Z=Le,oe=_e,ne=Ut,ie=zt,Be=os,Ce=Qt,Fe=Ne,He=Xt,je=Zt,xe=Pt,Ye=Ae,Ve=Oe;return o(),_("div",{class:"detail-item",onClick:i[7]||(i[7]=H=>G(s.value.id))},[t(Ve,null,{avatar:n(()=>[t(F,{round:"",size:30,src:s.value.user.avatar},null,8,["src"])]),header:n(()=>[t(le,{onClick:i[0]||(i[0]=Y(()=>{},["stop"])),class:"username-link",to:{name:"user",query:{s:s.value.user.username}}},{default:n(()=>[O(A(s.value.user.nickname),1)]),_:1},8,["to"]),m("span",eo," @"+A(s.value.user.username),1),s.value.is_top?(o(),z(Z,{key:0,class:"top-tag",type:"warning",size:"small",round:""},{default:n(()=>i[8]||(i[8]=[O(" 置顶 ")])),_:1})):u("",!0),s.value.visibility==a(ee).PRIVATE?(o(),z(Z,{key:1,class:"top-tag",type:"error",size:"small",round:""},{default:n(()=>i[9]||(i[9]=[O(" 私密 ")])),_:1})):u("",!0),s.value.visibility==a(ee).FRIEND?(o(),z(Z,{key:2,class:"top-tag",type:"info",size:"small",round:""},{default:n(()=>i[10]||(i[10]=[O(" 好友可见 ")])),_:1})):u("",!0)]),"header-extra":n(()=>[m("div",to,[t(ne,{placement:"bottom-end",trigger:"click",size:"small",options:E.value,onSelect:R},{default:n(()=>[t(oe,{quaternary:"",circle:""},{icon:n(()=>[t(a(K),null,{default:n(()=>[t(a(Yt))]),_:1})]),_:1})]),_:1},8,["options"])]),t(ie,{show:l.value,"onUpdate:show":i[1]||(i[1]=H=>l.value=H),"mask-closable":!1,preset:"dialog",title:"提示",content:"确定删除该泡泡动态吗？","positive-text":"确认","negative-text":"取消",onPositiveClick:y},null,8,["show"]),t(ie,{show:q.value,"onUpdate:show":i[2]||(i[2]=H=>q.value=H),"mask-closable":!1,preset:"dialog",title:"提示",content:"确定"+(s.value.is_lock?"解锁":"锁定")+"该泡泡动态吗？","positive-text":"确认","negative-text":"取消",onPositiveClick:Q},null,8,["show","content"]),t(ie,{show:D.value,"onUpdate:show":i[3]||(i[3]=H=>D.value=H),"mask-closable":!1,preset:"dialog",title:"提示",content:"确定"+(s.value.is_top?"取消置顶":"置顶")+"该泡泡动态吗？","positive-text":"确认","negative-text":"取消",onPositiveClick:X},null,8,["show","content"]),t(ie,{show:g.value,"onUpdate:show":i[4]||(i[4]=H=>g.value=H),"mask-closable":!1,preset:"dialog",title:"提示",content:"确定将该泡泡动态"+(s.value.is_essence?"取消亮点":"设为亮点")+"吗？","positive-text":"确认","negative-text":"取消",onPositiveClick:h},null,8,["show","content"]),t(ie,{show:$.value,"onUpdate:show":i[5]||(i[5]=H=>$.value=H),"mask-closable":!1,preset:"dialog",title:"提示",content:"确定将该泡泡动态可见度修改为"+(P.value==0?"公开":P.value==1?"私密":P.value==2?"好友可见":"关注可见")+"吗？","positive-text":"确认","negative-text":"取消",onPositiveClick:p},null,8,["show","content"]),t(Be,{show:C.value,user:w.value,onSuccess:I},null,8,["show","user"])]),footer:n(()=>[t(Ce,{attachments:s.value.attachments},null,8,["attachments"]),t(Ce,{attachments:s.value.charge_attachments,price:s.value.attachment_price},null,8,["attachments","price"]),t(Fe,{imgs:s.value.imgs},null,8,["imgs"]),t(He,{videos:s.value.videos,full:!0},null,8,["videos"]),t(je,{links:s.value.links},null,8,["links"]),m("div",no,[O(" 发布于 "+A(a(he)(s.value.created_on))+" ",1),s.value.ip_loc?(o(),_("span",ao,[t(xe,{vertical:""}),O(" "+A(s.value.ip_loc),1)])):u("",!0),!a(d).state.collapsedLeft&&s.value.created_on!=s.value.latest_replied_on?(o(),_("span",lo,[t(xe,{vertical:""}),O(" 最后回复 "+A(a(he)(s.value.latest_replied_on)),1)])):u("",!0)])]),action:n(()=>[m("div",io,[t(Ye,{justify:"space-between"},{default:n(()=>[m("div",{class:"opt-item hover",onClick:Y(T,["stop"])},[t(a(K),{size:"20",class:"opt-item-icon"},{default:n(()=>[L.value?u("",!0):(o(),z(a(Nt),{key:0})),L.value?(o(),z(a(Et),{key:1,color:"red"})):u("",!0)]),_:1}),O(" "+A(s.value.upvote_count),1)]),m("div",uo,[t(a(K),{size:"20",class:"opt-item-icon"},{default:n(()=>[t(a(Bt))]),_:1}),O(" "+A(s.value.comment_count),1)]),m("div",{class:"opt-item hover",onClick:Y(S,["stop"])},[t(a(K),{size:"20",class:"opt-item-icon"},{default:n(()=>[U.value?u("",!0):(o(),z(a(Ft),{key:0})),U.value?(o(),z(a(Ht),{key:1,color:"#ff7600"})):u("",!0)]),_:1}),O(" "+A(s.value.collection_count),1)]),m("div",{class:"opt-item hover",onClick:Y(B,["stop"])},[t(a(K),{size:"20",class:"opt-item-icon"},{default:n(()=>[t(a(jt))]),_:1}),O(" "+A(s.value.share_count),1)])]),_:1})])]),default:n(()=>[s.value.texts.length>0?(o(),_("div",so,[(o(!0),_(ve,null,fe(s.value.texts,H=>(o(),_("span",{key:H.id,class:"post-text",onClick:i[6]||(i[6]=Y(Ge=>W(Ge,s.value.id),["stop"])),innerHTML:a($e)(H.content).content},null,8,oo))),128))])):u("",!0)]),_:1})])}}}),ro={key:0,class:"detail-wrap"},po={key:1,class:"empty-wrap"},_o={key:0,class:"comment-opts-wrap"},mo={key:2},vo={key:0,class:"skeleton-wrap"},fo={key:1},ho={key:0,class:"empty-wrap"},go={key:0,class:"load-more-spinner"},yo={key:1,class:"load-more-spinner"},ko={key:2,class:"load-more-spinner"},wo={key:3,class:"load-more-spinner"},bo={key:4,class:"load-more-spinner"},$o={key:5,class:"load-more-spinner"},te=20,Co=ae({__name:"Post",setup(J){const j=es(),c=r({}),d=r(!1),k=r(!1),x=r([]),L=ce(()=>+j.query.id),U=r("default"),f=r(!0);let l={loading(){},loaded(){},complete(){},error(){}};const q=R=>{U.value=R,R==="default"&&(f.value=!0),G(l)},D=R=>{Pe({id:R}).then(y=>{c.value=y}).catch(y=>{})},g=()=>{c.value={id:0},d.value=!0,Pe({id:L.value}).then(R=>{d.value=!1,c.value=R,G(l)}).catch(R=>{d.value=!1})};let $=1;const M=r(!1),P=r([]),C=R=>{M.value||ke({id:c.value.id,style:"default",page:$,page_size:te}).then(y=>{R!==null&&(l=R),y.list.length<te?M.value=!0:$++,y.list.length>0&&($===1?P.value=y.list:P.value.push(...y.list),x.value=P.value),l.loaded(),k.value=!1}).catch(y=>{k.value=!1,l.error()})};let w=1,b=r(!1);const I=r([]),N=R=>{b.value||ke({id:c.value.id,style:"hots",page:w,page_size:te}).then(y=>{R!==null&&(l=R),y.list.length<te?b.value=!0:w++,y.list.length>0&&(w===1?I.value=y.list:I.value.push(...y.list),x.value=I.value),l.loaded(),k.value=!1}).catch(y=>{k.value=!1,l.error()})};let s=1,v=r(!1);const E=r([]),V=R=>{v.value||ke({id:c.value.id,style:"newest",page:s,page_size:te}).then(y=>{R!==null&&(l=R),y.list.length<te?v.value=!0:s++,y.list.length>0&&(s===1?E.value=y.list:E.value.push(...y.list),x.value=E.value),l.loaded(),k.value=!1}).catch(y=>{k.value=!1,l.error()})},G=R=>{L.value<1||(x.value.length===0&&(k.value=!0),U.value==="default"?(x.value=P.value,C(R)):U.value==="hots"?(x.value=I.value,N(R)):(x.value=E.value,V(R)),k.value=!1)},W=()=>{$=1,M.value=!1,P.value=[],w=1,b.value=!1,I.value=[],s=1,v.value=!1,E.value=[],G(l)};return Re(()=>{g()}),Ke(L,()=>{L.value>0&&j.name==="post"&&g()}),(R,y)=>{const Q=as,X=co,h=Ot,p=St,T=Rt,S=At,B=Lt,e=Zs,i=ts,F=Bs,le=Ae,Z=It;return o(),_("div",null,[t(Q,{title:"泡泡详情",back:!0}),t(Z,{class:"main-content-wrap",bordered:""},{default:n(()=>[t(T,null,{default:n(()=>[t(p,{show:d.value},{default:n(()=>[c.value.id>0?(o(),_("div",ro,[t(X,{post:c.value,onReload:D},null,8,["post"])])):(o(),_("div",po,[t(h,{size:"large",description:"暂无数据"})]))]),_:1},8,["show"])]),_:1}),c.value.id>0?(o(),_("div",_o,[t(B,{type:"bar","justify-content":"end",size:"small","tab-style":"margin-left: -24px;",animated:"","onUpdate:value":q},{prefix:n(()=>y[0]||(y[0]=[m("span",{class:"comment-title-item"},"评论",-1)])),default:n(()=>[t(S,{name:"default",tab:"推荐"}),t(S,{name:"hots",tab:"热门"}),t(S,{name:"newest",tab:"最新"})]),_:1})])):u("",!0),c.value.id>0?(o(),z(T,{key:1},{default:n(()=>[t(e,{lock:c.value.is_lock,"post-id":c.value.id,onPostSuccess:W},null,8,["lock","post-id"])]),_:1})):u("",!0),c.value.id>0?(o(),_("div",mo,[k.value?(o(),_("div",vo,[t(i,{num:5})])):(o(),_("div",fo,[x.value.length===0?(o(),_("div",ho,[t(h,{size:"large",description:"暂无评论，快来抢沙发"})])):u("",!0),(o(!0),_(ve,null,fe(x.value,oe=>(o(),z(T,{key:oe.id},{default:n(()=>[t(F,{comment:oe,postUserId:c.value.user_id,onReload:W},null,8,["comment","postUserId"])]),_:2},1024))),128))]))])):u("",!0),x.value.length>=te?(o(),z(le,{key:3,justify:"center"},{default:n(()=>[t(a(ls),{class:"load-more",slots:{complete:"没有更多数据了",error:"加载出错"},onInfinite:G},{spinner:n(()=>[f.value&&M.value?(o(),_("span",go)):u("",!0),!f.value&&a(b)?(o(),_("span",yo)):u("",!0),!f.value&&a(v)?(o(),_("span",ko)):u("",!0),f.value&&!M.value?(o(),_("span",wo,"加载评论")):u("",!0),!f.value&&!a(b)?(o(),_("span",bo,"加载评论")):u("",!0),!f.value&&!a(v)?(o(),_("span",$o,"加载评论")):u("",!0)]),_:1})]),_:1})):u("",!0)]),_:1})])}}}),on=pe(Co,[["__scopeId","data-v-776596be"]]);export{on as default};
