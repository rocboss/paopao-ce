import{_ as Pe,a as Ue}from"./post-item.vue_vue_type_style_index_0_lang-3baf8ba8.js";import{_ as ze}from"./post-skeleton-41befd31.js";import{C as Fe,_ as D,D as Te,E as Me,B as S,F as Se,G as Ie}from"./index-08d8af97.js";import{R as Q,K as J,S as X,b as Y,e as N,i as Z,T as qe,F as Ce,G as We,o as Be,O as Re,j as Ae,P as De,k as Ne,f as Ve,g as Ee,I as He,H as Oe}from"./naive-ui-62663ad7.js";import{d as V,r as l,o as u,Q as F,a1 as o,a as p,V as a,e as w,M as y,E as je,n as Ge,w as Le,j as Ke,c as h,_ as b,O as z,F as L,a4 as K}from"./@vue-e0e89260.js";import{_ as Qe}from"./main-nav.vue_vue_type_style_index_0_lang-18d4a8d3.js";import{u as Je}from"./vuex-473b3783.js";import{b as Xe}from"./vue-router-b8e3382f.js";import{i as Ye}from"./@vicons-6332ad63.js";import"./content-91ba374b.js";import"./paopao-video-player-aa5e8b3f.js";import"./formatTime-cdf4e6f1.js";import"./moment-2ab8298d.js";import"./copy-to-clipboard-1dd3075d.js";import"./toggle-selection-93f4ad84.js";import"./axios-4a70c6fc.js";/* empty css               */import"./seemly-76b7b838.js";import"./vueuc-59ca65c3.js";import"./evtd-b614532e.js";import"./@css-render-580d83ec.js";import"./vooks-a50491fd.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";const Ze={class:"whisper-wrap"},es={class:"whisper-line"},ss={class:"whisper-line send-wrap"},as=V({__name:"whisper",props:{show:{type:Boolean,default:!1},user:{}},emits:["success"],setup(I,{emit:x}){const d=I,c=l(""),t=l(!1),s=()=>{x("success")},_=()=>{t.value=!0,Fe({user_id:d.user.id,content:c.value}).then(r=>{window.$message.success("发送成功"),t.value=!1,c.value="",s()}).catch(r=>{t.value=!1})};return(r,g)=>{const i=Q,m=J,n=X,f=Y,$=N,P=Z;return u(),F(P,{show:r.show,"onUpdate:show":s,class:"whisper-card",preset:"card",size:"small",title:"私信","mask-closable":!1,bordered:!1,style:{width:"360px"}},{default:o(()=>[p("div",Ze,[a(n,{"show-icon":!1},{default:o(()=>[w(" 即将发送私信给: "),a(m,{style:{"max-width":"100%"}},{default:o(()=>[a(i,{type:"success"},{default:o(()=>[w(y(r.user.nickname)+"@"+y(r.user.username),1)]),_:1})]),_:1})]),_:1}),p("div",es,[a(f,{type:"textarea",placeholder:"请输入私信内容（请勿发送不和谐内容，否则将会被封号）",autosize:{minRows:5,maxRows:10},value:c.value,"onUpdate:value":g[0]||(g[0]=U=>c.value=U),maxlength:"200","show-count":""},null,8,["value"])]),p("div",ss,[a($,{strong:"",secondary:"",type:"primary",loading:t.value,onClick:_},{default:o(()=>[w(" 发送 ")]),_:1},8,["loading"])])])]),_:1},8,["show"])}}});const ts=D(as,[["__scopeId","data-v-0cbfe47c"]]),ns={class:"whisper-wrap"},os={class:"whisper-line"},ls={class:"whisper-line send-wrap"},is=V({__name:"whisper-add-friend",props:{show:{type:Boolean,default:!1},user:{}},emits:["success"],setup(I,{emit:x}){const d=I,c=l(""),t=l(!1),s=()=>{x("success")},_=()=>{t.value=!0,Te({user_id:d.user.id,greetings:c.value}).then(r=>{window.$message.success("发送成功"),t.value=!1,c.value="",s()}).catch(r=>{t.value=!1})};return(r,g)=>{const i=Q,m=J,n=X,f=Y,$=N,P=Z;return u(),F(P,{show:r.show,"onUpdate:show":s,class:"whisper-card",preset:"card",size:"small",title:"申请添加朋友","mask-closable":!1,bordered:!1,style:{width:"360px"}},{default:o(()=>[p("div",ns,[a(n,{"show-icon":!1},{default:o(()=>[w(" 发送添加朋友申请给: "),a(m,{style:{"max-width":"100%"}},{default:o(()=>[a(i,{type:"success"},{default:o(()=>[w(y(r.user.nickname)+"@"+y(r.user.username),1)]),_:1})]),_:1})]),_:1}),p("div",os,[a(f,{type:"textarea",placeholder:"请输入真挚的问候语",autosize:{minRows:5,maxRows:10},value:c.value,"onUpdate:value":g[0]||(g[0]=U=>c.value=U),maxlength:"120","show-count":""},null,8,["value"])]),p("div",ls,[a($,{strong:"",secondary:"",type:"primary",loading:t.value,onClick:_},{default:o(()=>[w(" 发送 ")]),_:1},8,["loading"])])])]),_:1},8,["show"])}}});const us=D(is,[["__scopeId","data-v-60be56a2"]]),cs={key:0,class:"profile-baseinfo"},rs={class:"avatar"},_s={class:"base-info"},ps={class:"username"},ds={class:"uid"},ms={key:0,class:"user-opts"},vs={key:0,class:"skeleton-wrap"},hs={key:1},gs={key:0,class:"empty-wrap"},fs={key:1},ws={key:2},ks={key:0,class:"pagination-wrap"},bs=V({__name:"User",setup(I){const x=qe(),d=Je(),c=Xe(),t=l(!1),s=je({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!0,status:1}),_=l(!1),r=l(!1),g=l(!1),i=l([]),m=l(c.query.username||""),n=l(+c.query.p||1),f=l("post"),$=l(+c.query.p||1),P=l(1),U=l(1),E=l(1),H=l(1),v=l(20),k=l(0),ee=()=>{switch(f.value){case"post":q();break;case"comment":C();break;case"highlight":W();break;case"media":B();break;case"star":R();break}},q=()=>{t.value=!0,S({username:m.value,style:"post",page:n.value,page_size:v.value}).then(e=>{t.value=!1,i.value=e.list||[],k.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{i.value=[],t.value=!1})},C=()=>{t.value=!0,S({username:m.value,style:"comment",page:n.value,page_size:v.value}).then(e=>{t.value=!1,i.value=e.list||[],k.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{i.value=[],t.value=!1})},W=()=>{t.value=!0,S({username:m.value,style:"highlight",page:n.value,page_size:v.value}).then(e=>{t.value=!1,i.value=e.list||[],k.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{i.value=[],t.value=!1})},B=()=>{t.value=!0,S({username:m.value,style:"media",page:n.value,page_size:v.value}).then(e=>{t.value=!1,i.value=e.list||[],k.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{i.value=[],t.value=!1})},R=()=>{t.value=!0,S({username:m.value,style:"star",page:n.value,page_size:v.value}).then(e=>{t.value=!1,i.value=e.list||[],k.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{i.value=[],t.value=!1})},se=e=>{switch(f.value=e,f.value){case"post":n.value=$.value,q();break;case"comment":n.value=P.value,C();break;case"highlight":n.value=U.value,W();break;case"media":n.value=E.value,B();break;case"star":n.value=H.value,R();break}},A=()=>{_.value=!0,Me({username:m.value}).then(e=>{_.value=!1,s.id=e.id,s.avatar=e.avatar,s.username=e.username,s.nickname=e.nickname,s.is_admin=e.is_admin,s.is_friend=e.is_friend,s.status=e.status,ee()}).catch(e=>{_.value=!1,console.log(e)})},ae=e=>{switch(n.value=e,f.value){case"post":$.value=e,q();break;case"comment":P.value=n.value,C();break;case"highlight":U.value=n.value,W();break;case"media":E.value=n.value,B();break;case"star":H.value=n.value,R();break}},te=()=>{r.value=!0},ne=()=>{g.value=!0},oe=()=>{r.value=!1},le=()=>{g.value=!1},ie=Ge(()=>{let e=[{label:"私信",key:"whisper"}];return d.state.userInfo.is_admin&&(s.status===1?e.push({label:"禁言",key:"banned"}):e.push({label:"解封",key:"deblocking"})),s.is_friend?e.push({label:"删除好友",key:"delete"}):e.push({label:"添加朋友",key:"requesting"}),e}),ue=e=>{switch(e){case"whisper":te();break;case"delete":ce();break;case"requesting":ne();break;case"banned":case"deblocking":re();break}},ce=()=>{x.warning({title:"删除好友",content:"将好友 “"+s.nickname+"” 删除，将同时删除 点赞/收藏 列表中关于该朋友的 “好友可见” 推文",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{_.value=!0,Se({user_id:s.id}).then(e=>{_.value=!1,s.is_friend=!1,q()}).catch(e=>{_.value=!1,console.log(e)})}})},re=()=>{x.warning({title:"警告",content:"确定对该用户进行"+(s.status===1?"禁言":"解封")+"处理吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{_.value=!0,Ie({id:s.id,status:s.status===1?2:1}).then(e=>{_.value=!1,A()}).catch(e=>{_.value=!1,console.log(e)})}})};return Le(()=>({path:c.path,query:c.query}),(e,O)=>{O.path==="/user"&&e.path==="/user"&&(m.value=c.query.username||"",A())}),Ke(()=>{A()}),(e,O)=>{const _e=Qe,pe=Be,j=Re,de=Ae,me=N,ve=De,he=ts,ge=Ne,T=Ve,fe=Ee,we=ze,ke=He,be=Pe,G=Oe,ye=Ue,xe=Ce,$e=We;return u(),h("div",null,[a(_e,{title:"用户详情"}),a(xe,{class:"main-content-wrap profile-wrap",bordered:""},{default:o(()=>[a(ge,{show:_.value},{default:o(()=>[s.id>0?(u(),h("div",cs,[p("div",rs,[a(pe,{size:"large",src:s.avatar},null,8,["src"])]),p("div",_s,[p("div",ps,[p("strong",null,y(s.nickname),1),p("span",null," @"+y(s.username),1),b(d).state.userInfo.id>0&&b(d).state.userInfo.username!=s.username&&s.is_friend?(u(),F(j,{key:0,class:"top-tag",type:"info",size:"small",round:""},{default:o(()=>[w(" 好友 ")]),_:1})):z("",!0),s.is_admin?(u(),F(j,{key:1,class:"top-tag",type:"error",size:"small",round:""},{default:o(()=>[w(" 管理员 ")]),_:1})):z("",!0)]),p("div",ds,"UID. "+y(s.id),1)]),b(d).state.userInfo.id>0&&b(d).state.userInfo.username!=s.username?(u(),h("div",ms,[a(ve,{placement:"bottom-end",trigger:"click",size:"small",options:ie.value,onSelect:ue},{default:o(()=>[a(me,{quaternary:"",circle:""},{icon:o(()=>[a(de,null,{default:o(()=>[a(b(Ye))]),_:1})]),_:1})]),_:1},8,["options"])])):z("",!0)])):z("",!0),a(he,{show:r.value,user:s,onSuccess:oe},null,8,["show","user"]),a(us,{show:g.value,user:s,onSuccess:le},null,8,["show","user"])]),_:1},8,["show"]),a(fe,{class:"profile-tabs-wrap",type:"line",animated:"","onUpdate:value":se},{default:o(()=>[a(T,{name:"post",tab:"泡泡"}),a(T,{name:"comment",tab:"评论"}),a(T,{name:"highlight",tab:"亮点"}),a(T,{name:"media",tab:"图文"}),a(T,{name:"star",tab:"喜欢"})]),_:1}),t.value?(u(),h("div",vs,[a(we,{num:v.value},null,8,["num"])])):(u(),h("div",hs,[i.value.length===0?(u(),h("div",gs,[a(ke,{size:"large",description:"暂无数据"})])):z("",!0),b(d).state.desktopModelShow?(u(),h("div",fs,[(u(!0),h(L,null,K(i.value,M=>(u(),F(G,{key:M.id},{default:o(()=>[a(be,{post:M},null,8,["post"])]),_:2},1024))),128))])):(u(),h("div",ws,[(u(!0),h(L,null,K(i.value,M=>(u(),F(G,{key:M.id},{default:o(()=>[a(ye,{post:M},null,8,["post"])]),_:2},1024))),128))]))]))]),_:1}),k.value>0?(u(),h("div",ks,[a($e,{page:n.value,"onUpdate:page":ae,"page-slot":b(d).state.collapsedRight?5:8,"page-count":k.value},null,8,["page","page-slot","page-count"])])):z("",!0)])}}});const Ys=D(bs,[["__scopeId","data-v-b67c9295"]]);export{Ys as default};
