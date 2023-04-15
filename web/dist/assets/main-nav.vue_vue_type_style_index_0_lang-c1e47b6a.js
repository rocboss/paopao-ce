import{cL as Be,cM as Ie,cN as He,bq as Oe,r as S,d as H,q as Le,cO as Se,y as C,u as Ce,j as Ve,bJ as je,S as Ue,c1 as Ne,cP as We,p as ae,bw as xe,bo as Re,h as i,cQ as Ye,b0 as Pe,b7 as Xe,cd as De,cR as qe,cS as Ke,cT as Ze,b as s,cU as ce,c as k,e as v,f as a,cV as Ge,aY as Je,x as de,cW as Qe,s as ze,t as $e,ba as Me,cX as et,A as Ae,cY as tt,cZ as nt,c_ as ot,w as P,b8 as rt,cG as it,k as at,c$ as st,m as lt,cB as Ee,a as Te,o as ct,z as W,br as ye,c7 as D,d0 as _e,B as Y,cz as dt,cc as ut,W as I,Y as Z,Z as oe,a2 as ht,a3 as X,a4 as L,a5 as A,a7 as re,a6 as ke,a9 as ft,aa as bt,ab as gt,$ as mt,ak as vt,d1 as wt,af as pt,K as yt,bP as _t}from"./index-db059f1c.js";let ie=0;const kt=typeof window<"u"&&window.matchMedia!==void 0,K=S(null);let R,q;function se(e){e.matches&&(K.value="dark")}function le(e){e.matches&&(K.value="light")}function xt(){R=window.matchMedia("(prefers-color-scheme: dark)"),q=window.matchMedia("(prefers-color-scheme: light)"),R.matches?K.value="dark":q.matches?K.value="light":K.value=null,R.addEventListener?(R.addEventListener("change",se),q.addEventListener("change",le)):R.addListener&&(R.addListener(se),q.addListener(le))}function zt(){"removeEventListener"in R?(R.removeEventListener("change",se),q.removeEventListener("change",le)):"removeListener"in R&&(R.removeListener(se),q.removeListener(le)),R=void 0,q=void 0}let Fe=!0;function $t(){return kt?(ie===0&&xt(),Fe&&(Fe=Ie())&&(He(()=>{ie+=1}),Oe(()=>{ie-=1,ie===0&&zt()})),Be(K)):Be(K)}const St=H({name:"NDrawerContent",inheritAttrs:!1,props:{blockScroll:Boolean,show:{type:Boolean,default:void 0},displayDirective:{type:String,required:!0},placement:{type:String,required:!0},contentStyle:[Object,String],nativeScrollbar:{type:Boolean,required:!0},scrollbarProps:Object,trapFocus:{type:Boolean,default:!0},autoFocus:{type:Boolean,default:!0},showMask:{type:[Boolean,String],required:!0},resizable:Boolean,onClickoutside:Function,onAfterLeave:Function,onAfterEnter:Function,onEsc:Function},setup(e){const t=S(!!e.show),n=S(null),d=Le(Se);let c=0,m="",u=null;const w=S(!1),p=S(!1),g=C(()=>e.placement==="top"||e.placement==="bottom"),{mergedClsPrefixRef:b,mergedRtlRef:x}=Ce(e),h=Ve("Drawer",x,b),_=r=>{p.value=!0,c=g.value?r.clientY:r.clientX,m=document.body.style.cursor,document.body.style.cursor=g.value?"ns-resize":"ew-resize",document.body.addEventListener("mousemove",M),document.body.addEventListener("mouseleave",B),document.body.addEventListener("mouseup",E)},y=()=>{u!==null&&(window.clearTimeout(u),u=null),p.value?w.value=!0:u=window.setTimeout(()=>{w.value=!0},300)},V=()=>{u!==null&&(window.clearTimeout(u),u=null),w.value=!1},{doUpdateHeight:j,doUpdateWidth:U}=d,M=r=>{var o,T;if(p.value)if(g.value){let z=((o=n.value)===null||o===void 0?void 0:o.offsetHeight)||0;const $=c-r.clientY;z+=e.placement==="bottom"?$:-$,j(z),c=r.clientY}else{let z=((T=n.value)===null||T===void 0?void 0:T.offsetWidth)||0;const $=c-r.clientX;z+=e.placement==="right"?$:-$,U(z),c=r.clientX}},E=()=>{p.value&&(c=0,p.value=!1,document.body.style.cursor=m,document.body.removeEventListener("mousemove",M),document.body.removeEventListener("mouseup",E),document.body.removeEventListener("mouseleave",B))},B=E;je(()=>{e.show&&(t.value=!0)}),Ue(()=>e.show,r=>{r||E()}),Oe(()=>{E()});const l=C(()=>{const{show:r}=e,o=[[Re,r]];return e.showMask||o.push([Ne,e.onClickoutside,void 0,{capture:!0}]),o});function f(){var r;t.value=!1,(r=e.onAfterLeave)===null||r===void 0||r.call(e)}return We(C(()=>e.blockScroll&&t.value)),ae(qe,n),ae(Ke,null),ae(Ze,null),{bodyRef:n,rtlEnabled:h,mergedClsPrefix:d.mergedClsPrefixRef,isMounted:d.isMountedRef,mergedTheme:d.mergedThemeRef,displayed:t,transitionName:C(()=>({right:"slide-in-from-right-transition",left:"slide-in-from-left-transition",top:"slide-in-from-top-transition",bottom:"slide-in-from-bottom-transition"})[e.placement]),handleAfterLeave:f,bodyDirectives:l,handleMousedownResizeTrigger:_,handleMouseenterResizeTrigger:y,handleMouseleaveResizeTrigger:V,isDragging:p,isHoverOnResizeTrigger:w}},render(){const{$slots:e,mergedClsPrefix:t}=this;return this.displayDirective==="show"||this.displayed||this.show?xe(i("div",{role:"none"},i(Ye,{disabled:!this.showMask||!this.trapFocus,active:this.show,autoFocus:this.autoFocus,onEsc:this.onEsc},{default:()=>i(Pe,{name:this.transitionName,appear:this.isMounted,onAfterEnter:this.onAfterEnter,onAfterLeave:this.handleAfterLeave},{default:()=>xe(i("div",Xe(this.$attrs,{role:"dialog",ref:"bodyRef","aria-modal":"true",class:[`${t}-drawer`,this.rtlEnabled&&`${t}-drawer--rtl`,`${t}-drawer--${this.placement}-placement`,this.isDragging&&`${t}-drawer--unselectable`,this.nativeScrollbar&&`${t}-drawer--native-scrollbar`]}),[this.resizable?i("div",{class:[`${t}-drawer__resize-trigger`,(this.isDragging||this.isHoverOnResizeTrigger)&&`${t}-drawer__resize-trigger--hover`],onMouseenter:this.handleMouseenterResizeTrigger,onMouseleave:this.handleMouseleaveResizeTrigger,onMousedown:this.handleMousedownResizeTrigger}):null,this.nativeScrollbar?i("div",{class:`${t}-drawer-content-wrapper`,style:this.contentStyle,role:"none"},e):i(De,Object.assign({},this.scrollbarProps,{contentStyle:this.contentStyle,contentClass:`${t}-drawer-content-wrapper`,theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar}),e)]),this.bodyDirectives)})})),[[Re,this.displayDirective==="if"||this.displayed||this.show]]):null}}),{cubicBezierEaseIn:Ct,cubicBezierEaseOut:Bt}=ce;function Rt({duration:e="0.3s",leaveDuration:t="0.2s",name:n="slide-in-from-right"}={}){return[s(`&.${n}-transition-leave-active`,{transition:`transform ${t} ${Ct}`}),s(`&.${n}-transition-enter-active`,{transition:`transform ${e} ${Bt}`}),s(`&.${n}-transition-enter-to`,{transform:"translateX(0)"}),s(`&.${n}-transition-enter-from`,{transform:"translateX(100%)"}),s(`&.${n}-transition-leave-from`,{transform:"translateX(0)"}),s(`&.${n}-transition-leave-to`,{transform:"translateX(100%)"})]}const{cubicBezierEaseIn:Mt,cubicBezierEaseOut:Et}=ce;function Tt({duration:e="0.3s",leaveDuration:t="0.2s",name:n="slide-in-from-left"}={}){return[s(`&.${n}-transition-leave-active`,{transition:`transform ${t} ${Mt}`}),s(`&.${n}-transition-enter-active`,{transition:`transform ${e} ${Et}`}),s(`&.${n}-transition-enter-to`,{transform:"translateX(0)"}),s(`&.${n}-transition-enter-from`,{transform:"translateX(-100%)"}),s(`&.${n}-transition-leave-from`,{transform:"translateX(0)"}),s(`&.${n}-transition-leave-to`,{transform:"translateX(-100%)"})]}const{cubicBezierEaseIn:Ft,cubicBezierEaseOut:Ot}=ce;function Lt({duration:e="0.3s",leaveDuration:t="0.2s",name:n="slide-in-from-top"}={}){return[s(`&.${n}-transition-leave-active`,{transition:`transform ${t} ${Ft}`}),s(`&.${n}-transition-enter-active`,{transition:`transform ${e} ${Ot}`}),s(`&.${n}-transition-enter-to`,{transform:"translateY(0)"}),s(`&.${n}-transition-enter-from`,{transform:"translateY(-100%)"}),s(`&.${n}-transition-leave-from`,{transform:"translateY(0)"}),s(`&.${n}-transition-leave-to`,{transform:"translateY(-100%)"})]}const{cubicBezierEaseIn:Pt,cubicBezierEaseOut:Dt}=ce;function At({duration:e="0.3s",leaveDuration:t="0.2s",name:n="slide-in-from-bottom"}={}){return[s(`&.${n}-transition-leave-active`,{transition:`transform ${t} ${Pt}`}),s(`&.${n}-transition-enter-active`,{transition:`transform ${e} ${Dt}`}),s(`&.${n}-transition-enter-to`,{transform:"translateY(0)"}),s(`&.${n}-transition-enter-from`,{transform:"translateY(100%)"}),s(`&.${n}-transition-leave-from`,{transform:"translateY(0)"}),s(`&.${n}-transition-leave-to`,{transform:"translateY(100%)"})]}const It=s([k("drawer",`
 word-break: break-word;
 line-height: var(--n-line-height);
 position: absolute;
 pointer-events: all;
 box-shadow: var(--n-box-shadow);
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 background-color: var(--n-color);
 color: var(--n-text-color);
 box-sizing: border-box;
 `,[Rt(),Tt(),Lt(),At(),v("unselectable",`
 user-select: none; 
 -webkit-user-select: none;
 `),v("native-scrollbar",[k("drawer-content-wrapper",`
 overflow: auto;
 height: 100%;
 `)]),a("resize-trigger",`
 position: absolute;
 background-color: #0000;
 transition: background-color .3s var(--n-bezier);
 `,[v("hover",`
 background-color: var(--n-resize-trigger-color-hover);
 `)]),k("drawer-content-wrapper",`
 box-sizing: border-box;
 `),k("drawer-content",`
 height: 100%;
 display: flex;
 flex-direction: column;
 `,[v("native-scrollbar",[k("drawer-body-content-wrapper",`
 height: 100%;
 overflow: auto;
 `)]),k("drawer-body",`
 flex: 1 0 0;
 overflow: hidden;
 `),k("drawer-body-content-wrapper",`
 box-sizing: border-box;
 padding: var(--n-body-padding);
 `),k("drawer-header",`
 font-weight: var(--n-title-font-weight);
 line-height: 1;
 font-size: var(--n-title-font-size);
 color: var(--n-title-text-color);
 padding: var(--n-header-padding);
 transition: border .3s var(--n-bezier);
 border-bottom: 1px solid var(--n-divider-color);
 border-bottom: var(--n-header-border-bottom);
 display: flex;
 justify-content: space-between;
 align-items: center;
 `,[a("close",`
 margin-left: 6px;
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `)]),k("drawer-footer",`
 display: flex;
 justify-content: flex-end;
 border-top: var(--n-footer-border-top);
 transition: border .3s var(--n-bezier);
 padding: var(--n-footer-padding);
 `)]),v("right-placement",`
 top: 0;
 bottom: 0;
 right: 0;
 `,[a("resize-trigger",`
 width: 3px;
 height: 100%;
 top: 0;
 left: 0;
 transform: translateX(-1.5px);
 cursor: ew-resize;
 `)]),v("left-placement",`
 top: 0;
 bottom: 0;
 left: 0;
 `,[a("resize-trigger",`
 width: 3px;
 height: 100%;
 top: 0;
 right: 0;
 transform: translateX(1.5px);
 cursor: ew-resize;
 `)]),v("top-placement",`
 top: 0;
 left: 0;
 right: 0;
 `,[a("resize-trigger",`
 width: 100%;
 height: 3px;
 bottom: 0;
 left: 0;
 transform: translateY(1.5px);
 cursor: ns-resize;
 `)]),v("bottom-placement",`
 left: 0;
 bottom: 0;
 right: 0;
 `,[a("resize-trigger",`
 width: 100%;
 height: 3px;
 top: 0;
 left: 0;
 transform: translateY(-1.5px);
 cursor: ns-resize;
 `)])]),s("body",[s(">",[k("drawer-container",{position:"fixed"})])]),k("drawer-container",`
 position: relative;
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 pointer-events: none;
 `,[s("> *",{pointerEvents:"all"})]),k("drawer-mask",`
 background-color: rgba(0, 0, 0, .3);
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `,[v("invisible",`
 background-color: rgba(0, 0, 0, 0)
 `),Ge({enterDuration:"0.2s",leaveDuration:"0.2s",enterCubicBezier:"var(--n-bezier-in)",leaveCubicBezier:"var(--n-bezier-out)"})])]),Ht=Object.assign(Object.assign({},de.props),{show:Boolean,width:[Number,String],height:[Number,String],placement:{type:String,default:"right"},maskClosable:{type:Boolean,default:!0},showMask:{type:[Boolean,String],default:!0},to:[String,Object],displayDirective:{type:String,default:"if"},nativeScrollbar:{type:Boolean,default:!0},zIndex:Number,onMaskClick:Function,scrollbarProps:Object,contentStyle:[Object,String],trapFocus:{type:Boolean,default:!0},onEsc:Function,autoFocus:{type:Boolean,default:!0},closeOnEsc:{type:Boolean,default:!0},blockScroll:{type:Boolean,default:!0},resizable:Boolean,defaultWidth:{type:[Number,String],default:251},defaultHeight:{type:[Number,String],default:251},onUpdateWidth:[Function,Array],onUpdateHeight:[Function,Array],"onUpdate:width":[Function,Array],"onUpdate:height":[Function,Array],"onUpdate:show":[Function,Array],onUpdateShow:[Function,Array],onAfterEnter:Function,onAfterLeave:Function,drawerStyle:[String,Object],drawerClass:String,target:null,onShow:Function,onHide:Function}),Vt=H({name:"Drawer",inheritAttrs:!1,props:Ht,setup(e){const{mergedClsPrefixRef:t,namespaceRef:n,inlineThemeDisabled:d}=Ce(e),c=Je(),m=de("Drawer","-drawer",It,Qe,e,t),u=S(e.defaultWidth),w=S(e.defaultHeight),p=ze($e(e,"width"),u),g=ze($e(e,"height"),w),b=C(()=>{const{placement:l}=e;return l==="top"||l==="bottom"?"":Me(p.value)}),x=C(()=>{const{placement:l}=e;return l==="left"||l==="right"?"":Me(g.value)}),h=l=>{const{onUpdateWidth:f,"onUpdate:width":r}=e;f&&P(f,l),r&&P(r,l),u.value=l},_=l=>{const{onUpdateHeight:f,"onUpdate:width":r}=e;f&&P(f,l),r&&P(r,l),w.value=l},y=C(()=>[{width:b.value,height:x.value},e.drawerStyle||""]);function V(l){const{onMaskClick:f,maskClosable:r}=e;r&&M(!1),f&&f(l)}const j=et();function U(l){var f;(f=e.onEsc)===null||f===void 0||f.call(e),e.show&&e.closeOnEsc&&ot(l)&&!j.value&&M(!1)}function M(l){const{onHide:f,onUpdateShow:r,"onUpdate:show":o}=e;r&&P(r,l),o&&P(o,l),f&&!l&&P(f,l)}ae(Se,{isMountedRef:c,mergedThemeRef:m,mergedClsPrefixRef:t,doUpdateShow:M,doUpdateHeight:_,doUpdateWidth:h});const E=C(()=>{const{common:{cubicBezierEaseInOut:l,cubicBezierEaseIn:f,cubicBezierEaseOut:r},self:{color:o,textColor:T,boxShadow:z,lineHeight:$,headerPadding:G,footerPadding:J,bodyPadding:ue,titleFontSize:he,titleTextColor:fe,titleFontWeight:be,headerBorderBottom:F,footerBorderTop:ge,closeIconColor:me,closeIconColorHover:O,closeIconColorPressed:N,closeColorHover:ve,closeColorPressed:we,closeIconSize:pe,closeSize:Q,closeBorderRadius:ee,resizableTriggerColorHover:te}}=m.value;return{"--n-line-height":$,"--n-color":o,"--n-text-color":T,"--n-box-shadow":z,"--n-bezier":l,"--n-bezier-out":r,"--n-bezier-in":f,"--n-header-padding":G,"--n-body-padding":ue,"--n-footer-padding":J,"--n-title-text-color":fe,"--n-title-font-size":he,"--n-title-font-weight":be,"--n-header-border-bottom":F,"--n-footer-border-top":ge,"--n-close-icon-color":me,"--n-close-icon-color-hover":O,"--n-close-icon-color-pressed":N,"--n-close-size":Q,"--n-close-color-hover":ve,"--n-close-color-pressed":we,"--n-close-icon-size":pe,"--n-close-border-radius":ee,"--n-resize-trigger-color-hover":te}}),B=d?Ae("drawer",void 0,E,e):void 0;return{mergedClsPrefix:t,namespace:n,mergedBodyStyle:y,handleMaskClick:V,handleEsc:U,mergedTheme:m,cssVars:d?void 0:E,themeClass:B==null?void 0:B.themeClass,onRender:B==null?void 0:B.onRender,isMounted:c}},render(){const{mergedClsPrefix:e}=this;return i(nt,{to:this.to,show:this.show},{default:()=>{var t;return(t=this.onRender)===null||t===void 0||t.call(this),xe(i("div",{class:[`${e}-drawer-container`,this.namespace,this.themeClass],style:this.cssVars,role:"none"},this.showMask?i(Pe,{name:"fade-in-transition",appear:this.isMounted},{default:()=>this.show?i("div",{"aria-hidden":!0,class:[`${e}-drawer-mask`,this.showMask==="transparent"&&`${e}-drawer-mask--invisible`],onClick:this.handleMaskClick}):null}):null,i(St,Object.assign({},this.$attrs,{class:[this.drawerClass,this.$attrs.class],style:[this.mergedBodyStyle,this.$attrs.style],blockScroll:this.blockScroll,contentStyle:this.contentStyle,placement:this.placement,scrollbarProps:this.scrollbarProps,show:this.show,displayDirective:this.displayDirective,nativeScrollbar:this.nativeScrollbar,onAfterEnter:this.onAfterEnter,onAfterLeave:this.onAfterLeave,trapFocus:this.trapFocus,autoFocus:this.autoFocus,resizable:this.resizable,showMask:this.showMask,onEsc:this.handleEsc,onClickoutside:this.handleMaskClick}),this.$slots)),[[tt,{zIndex:this.zIndex,enabled:this.show}]])}})}}),jt={title:{type:String},headerStyle:[Object,String],footerStyle:[Object,String],bodyStyle:[Object,String],bodyContentStyle:[Object,String],nativeScrollbar:{type:Boolean,default:!0},scrollbarProps:Object,closable:Boolean},Ut=H({name:"DrawerContent",props:jt,setup(){const e=Le(Se,null);e||rt("drawer-content","`n-drawer-content` must be placed inside `n-drawer`.");const{doUpdateShow:t}=e;function n(){t(!1)}return{handleCloseClick:n,mergedTheme:e.mergedThemeRef,mergedClsPrefix:e.mergedClsPrefixRef}},render(){const{title:e,mergedClsPrefix:t,nativeScrollbar:n,mergedTheme:d,bodyStyle:c,bodyContentStyle:m,headerStyle:u,footerStyle:w,scrollbarProps:p,closable:g,$slots:b}=this;return i("div",{role:"none",class:[`${t}-drawer-content`,n&&`${t}-drawer-content--native-scrollbar`]},b.header||e||g?i("div",{class:`${t}-drawer-header`,style:u,role:"none"},i("div",{class:`${t}-drawer-header__main`,role:"heading","aria-level":"1"},b.header!==void 0?b.header():e),g&&i(it,{onClick:this.handleCloseClick,clsPrefix:t,class:`${t}-drawer-header__close`,absolute:!0})):null,n?i("div",{class:`${t}-drawer-body`,style:c,role:"none"},i("div",{class:`${t}-drawer-body-content-wrapper`,style:m,role:"none"},b)):i(De,Object.assign({themeOverrides:d.peerOverrides.Scrollbar,theme:d.peers.Scrollbar},p,{class:`${t}-drawer-body`,contentClass:`${t}-drawer-body-content-wrapper`,contentStyle:m}),b),b.footer?i("div",{class:`${t}-drawer-footer`,style:w,role:"none"},b.footer()):null)}}),Nt=e=>{const{primaryColor:t,opacityDisabled:n,borderRadius:d,textColor3:c}=e,m="rgba(0, 0, 0, .14)";return Object.assign(Object.assign({},st),{iconColor:c,textColor:"white",loadingColor:t,opacityDisabled:n,railColor:m,railColorActive:t,buttonBoxShadow:"0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)",buttonColor:"#FFF",railBorderRadiusSmall:d,railBorderRadiusMedium:d,railBorderRadiusLarge:d,buttonBorderRadiusSmall:d,buttonBorderRadiusMedium:d,buttonBorderRadiusLarge:d,boxShadowFocus:`0 0 0 2px ${lt(t,{alpha:.2})}`})},Wt={name:"Switch",common:at,self:Nt},Yt=Wt,Xt=k("switch",`
 height: var(--n-height);
 min-width: var(--n-width);
 vertical-align: middle;
 user-select: none;
 -webkit-user-select: none;
 display: inline-flex;
 outline: none;
 justify-content: center;
 align-items: center;
`,[a("children-placeholder",`
 height: var(--n-rail-height);
 display: flex;
 flex-direction: column;
 overflow: hidden;
 pointer-events: none;
 visibility: hidden;
 `),a("rail-placeholder",`
 display: flex;
 flex-wrap: none;
 `),a("button-placeholder",`
 width: calc(1.75 * var(--n-rail-height));
 height: var(--n-rail-height);
 `),k("base-loading",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 font-size: calc(var(--n-button-width) - 4px);
 color: var(--n-loading-color);
 transition: color .3s var(--n-bezier);
 `,[Ee({left:"50%",top:"50%",originalTransform:"translateX(-50%) translateY(-50%)"})]),a("checked, unchecked",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 box-sizing: border-box;
 position: absolute;
 white-space: nowrap;
 top: 0;
 bottom: 0;
 display: flex;
 align-items: center;
 line-height: 1;
 `),a("checked",`
 right: 0;
 padding-right: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),a("unchecked",`
 left: 0;
 justify-content: flex-end;
 padding-left: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),s("&:focus",[a("rail",`
 box-shadow: var(--n-box-shadow-focus);
 `)]),v("round",[a("rail","border-radius: calc(var(--n-rail-height) / 2);",[a("button","border-radius: calc(var(--n-button-height) / 2);")])]),Te("disabled",[Te("icon",[v("rubber-band",[v("pressed",[a("rail",[a("button","max-width: var(--n-button-width-pressed);")])]),a("rail",[s("&:active",[a("button","max-width: var(--n-button-width-pressed);")])]),v("active",[v("pressed",[a("rail",[a("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])]),a("rail",[s("&:active",[a("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])])])])])]),v("active",[a("rail",[a("button","left: calc(100% - var(--n-button-width) - var(--n-offset))")])]),a("rail",`
 overflow: hidden;
 height: var(--n-rail-height);
 min-width: var(--n-rail-width);
 border-radius: var(--n-rail-border-radius);
 cursor: pointer;
 position: relative;
 transition:
 opacity .3s var(--n-bezier),
 background .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-rail-color);
 `,[a("button-icon",`
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 font-size: calc(var(--n-button-height) - 4px);
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 display: flex;
 justify-content: center;
 align-items: center;
 line-height: 1;
 `,[Ee()]),a("button",`
 align-items: center; 
 top: var(--n-offset);
 left: var(--n-offset);
 height: var(--n-button-height);
 width: var(--n-button-width-pressed);
 max-width: var(--n-button-width);
 border-radius: var(--n-button-border-radius);
 background-color: var(--n-button-color);
 box-shadow: var(--n-button-box-shadow);
 box-sizing: border-box;
 cursor: inherit;
 content: "";
 position: absolute;
 transition:
 background-color .3s var(--n-bezier),
 left .3s var(--n-bezier),
 opacity .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 `)]),v("active",[a("rail","background-color: var(--n-rail-color-active);")]),v("loading",[a("rail",`
 cursor: wait;
 `)]),v("disabled",[a("rail",`
 cursor: not-allowed;
 opacity: .5;
 `)])]),qt=Object.assign(Object.assign({},de.props),{size:{type:String,default:"medium"},value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,rubberBand:{type:Boolean,default:!0},onChange:[Function,Array]});let ne;const Kt=H({name:"Switch",props:qt,setup(e){ne===void 0&&(typeof CSS<"u"?typeof CSS.supports<"u"?ne=CSS.supports("width","max(1px)"):ne=!1:ne=!0);const{mergedClsPrefixRef:t,inlineThemeDisabled:n}=Ce(e),d=de("Switch","-switch",Xt,Yt,e,t),c=ct(e),{mergedSizeRef:m,mergedDisabledRef:u}=c,w=S(e.defaultValue),p=$e(e,"value"),g=ze(p,w),b=C(()=>g.value===e.checkedValue),x=S(!1),h=S(!1),_=C(()=>{const{railStyle:o}=e;if(o)return o({focused:h.value,checked:b.value})});function y(o){const{"onUpdate:value":T,onChange:z,onUpdateValue:$}=e,{nTriggerFormInput:G,nTriggerFormChange:J}=c;T&&P(T,o),$&&P($,o),z&&P(z,o),w.value=o,G(),J()}function V(){const{nTriggerFormFocus:o}=c;o()}function j(){const{nTriggerFormBlur:o}=c;o()}function U(){e.loading||u.value||(g.value!==e.checkedValue?y(e.checkedValue):y(e.uncheckedValue))}function M(){h.value=!0,V()}function E(){h.value=!1,j(),x.value=!1}function B(o){e.loading||u.value||o.key===" "&&(g.value!==e.checkedValue?y(e.checkedValue):y(e.uncheckedValue),x.value=!1)}function l(o){e.loading||u.value||o.key===" "&&(o.preventDefault(),x.value=!0)}const f=C(()=>{const{value:o}=m,{self:{opacityDisabled:T,railColor:z,railColorActive:$,buttonBoxShadow:G,buttonColor:J,boxShadowFocus:ue,loadingColor:he,textColor:fe,iconColor:be,[W("buttonHeight",o)]:F,[W("buttonWidth",o)]:ge,[W("buttonWidthPressed",o)]:me,[W("railHeight",o)]:O,[W("railWidth",o)]:N,[W("railBorderRadius",o)]:ve,[W("buttonBorderRadius",o)]:we},common:{cubicBezierEaseInOut:pe}}=d.value;let Q,ee,te;return ne?(Q=`calc((${O} - ${F}) / 2)`,ee=`max(${O}, ${F})`,te=`max(${N}, calc(${N} + ${F} - ${O}))`):(Q=ye((D(O)-D(F))/2),ee=ye(Math.max(D(O),D(F))),te=D(O)>D(F)?N:ye(D(N)+D(F)-D(O))),{"--n-bezier":pe,"--n-button-border-radius":we,"--n-button-box-shadow":G,"--n-button-color":J,"--n-button-width":ge,"--n-button-width-pressed":me,"--n-button-height":F,"--n-height":ee,"--n-offset":Q,"--n-opacity-disabled":T,"--n-rail-border-radius":ve,"--n-rail-color":z,"--n-rail-color-active":$,"--n-rail-height":O,"--n-rail-width":N,"--n-width":te,"--n-box-shadow-focus":ue,"--n-loading-color":he,"--n-text-color":fe,"--n-icon-color":be}}),r=n?Ae("switch",C(()=>m.value[0]),f,e):void 0;return{handleClick:U,handleBlur:E,handleFocus:M,handleKeyup:B,handleKeydown:l,mergedRailStyle:_,pressed:x,mergedClsPrefix:t,mergedValue:g,checked:b,mergedDisabled:u,cssVars:n?void 0:f,themeClass:r==null?void 0:r.themeClass,onRender:r==null?void 0:r.onRender}},render(){const{mergedClsPrefix:e,mergedDisabled:t,checked:n,mergedRailStyle:d,onRender:c,$slots:m}=this;c==null||c();const{checked:u,unchecked:w,icon:p,"checked-icon":g,"unchecked-icon":b}=m,x=!(_e(p)&&_e(g)&&_e(b));return i("div",{role:"switch","aria-checked":n,class:[`${e}-switch`,this.themeClass,x&&`${e}-switch--icon`,n&&`${e}-switch--active`,t&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`,this.rubberBand&&`${e}-switch--rubber-band`],tabindex:this.mergedDisabled?void 0:0,style:this.cssVars,onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},i("div",{class:`${e}-switch__rail`,"aria-hidden":"true",style:d},Y(u,h=>Y(w,_=>h||_?i("div",{"aria-hidden":!0,class:`${e}-switch__children-placeholder`},i("div",{class:`${e}-switch__rail-placeholder`},i("div",{class:`${e}-switch__button-placeholder`}),h),i("div",{class:`${e}-switch__rail-placeholder`},i("div",{class:`${e}-switch__button-placeholder`}),_)):null)),i("div",{class:`${e}-switch__button`},Y(p,h=>Y(g,_=>Y(b,y=>i(dt,null,{default:()=>this.loading?i(ut,{key:"loading",clsPrefix:e,strokeWidth:20}):this.checked&&(_||h)?i("div",{class:`${e}-switch__button-icon`,key:_?"checked-icon":"icon"},_||h):!this.checked&&(y||h)?i("div",{class:`${e}-switch__button-icon`,key:y?"unchecked-icon":"icon"},y||h):null})))),Y(u,h=>h&&i("div",{key:"checked",class:`${e}-switch__checked`},h)),Y(w,h=>h&&i("div",{key:"unchecked",class:`${e}-switch__unchecked`},h)))))}}),Zt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},Gt=oe("path",{d:"M14.71 6.71a.996.996 0 0 0-1.41 0L8.71 11.3a.996.996 0 0 0 0 1.41l4.59 4.59a.996.996 0 1 0 1.41-1.41L10.83 12l3.88-3.88c.39-.39.38-1.03 0-1.41z",fill:"currentColor"},null,-1),Jt=[Gt],Qt=H({name:"ChevronLeftRound",render:function(t,n){return I(),Z("svg",Zt,Jt)}}),en={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},tn=oe("path",{d:"M9.37 5.51A7.35 7.35 0 0 0 9.1 7.5c0 4.08 3.32 7.4 7.4 7.4c.68 0 1.35-.09 1.99-.27A7.014 7.014 0 0 1 12 19c-3.86 0-7-3.14-7-7c0-2.93 1.81-5.45 4.37-6.49zM12 3a9 9 0 1 0 9 9c0-.46-.04-.92-.1-1.36a5.389 5.389 0 0 1-4.4 2.26a5.403 5.403 0 0 1-3.14-9.8c-.44-.06-.9-.1-1.36-.1z",fill:"currentColor"},null,-1),nn=[tn],on=H({name:"DarkModeOutlined",render:function(t,n){return I(),Z("svg",en,nn)}}),rn={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},an=oe("path",{d:"M2 17c0 .55.45 1 1 1h18c.55 0 1-.45 1-1s-.45-1-1-1H3c-.55 0-1 .45-1 1zm0-5c0 .55.45 1 1 1h18c.55 0 1-.45 1-1s-.45-1-1-1H3c-.55 0-1 .45-1 1zm0-5c0 .55.45 1 1 1h18c.55 0 1-.45 1-1s-.45-1-1-1H3c-.55 0-1 .45-1 1z",fill:"currentColor"},null,-1),sn=[an],ln=H({name:"DehazeRound",render:function(t,n){return I(),Z("svg",rn,sn)}}),cn={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},dn=oe("path",{d:"M12 9c1.65 0 3 1.35 3 3s-1.35 3-3 3s-3-1.35-3-3s1.35-3 3-3m0-2c-2.76 0-5 2.24-5 5s2.24 5 5 5s5-2.24 5-5s-2.24-5-5-5zM2 13h2c.55 0 1-.45 1-1s-.45-1-1-1H2c-.55 0-1 .45-1 1s.45 1 1 1zm18 0h2c.55 0 1-.45 1-1s-.45-1-1-1h-2c-.55 0-1 .45-1 1s.45 1 1 1zM11 2v2c0 .55.45 1 1 1s1-.45 1-1V2c0-.55-.45-1-1-1s-1 .45-1 1zm0 18v2c0 .55.45 1 1 1s1-.45 1-1v-2c0-.55-.45-1-1-1s-1 .45-1 1zM5.99 4.58a.996.996 0 0 0-1.41 0a.996.996 0 0 0 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0s.39-1.03 0-1.41L5.99 4.58zm12.37 12.37a.996.996 0 0 0-1.41 0a.996.996 0 0 0 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0a.996.996 0 0 0 0-1.41l-1.06-1.06zm1.06-10.96a.996.996 0 0 0 0-1.41a.996.996 0 0 0-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06zM7.05 18.36a.996.996 0 0 0 0-1.41a.996.996 0 0 0-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06z",fill:"currentColor"},null,-1),un=[dn],hn=H({name:"LightModeOutlined",render:function(t,n){return I(),Z("svg",cn,un)}}),fn={key:0},bn={class:"navbar"},mn=H({__name:"main-nav",props:{title:{default:""},back:{type:Boolean,default:!1},theme:{type:Boolean,default:!0}},setup(e){const t=e,n=mt(),d=vt(),c=S(!1),m=S("left"),u=g=>{g?(localStorage.setItem("PAOPAO_THEME","dark"),n.commit("triggerTheme","dark")):(localStorage.setItem("PAOPAO_THEME","light"),n.commit("triggerTheme","light"))},w=()=>{window.history.length<=1?d.push({path:"/"}):d.go(-1)},p=()=>{c.value=!0};return ht(()=>{localStorage.getItem("PAOPAO_THEME")||u($t()==="dark")}),(g,b)=>{const x=wt,h=Ut,_=Vt,y=pt,V=yt,j=Kt,U=_t;return I(),Z(gt,null,[X(n).state.drawerModelShow?(I(),Z("div",fn,[L(_,{show:c.value,"onUpdate:show":b[0]||(b[0]=M=>c.value=M),placement:m.value,resizable:""},{default:A(()=>[L(h,null,{default:A(()=>[L(x)]),_:1})]),_:1},8,["show","placement"])])):re("",!0),L(U,{size:"small",bordered:!0,class:"nav-title-card"},{header:A(()=>[oe("div",bn,[X(n).state.drawerModelShow&&!e.back?(I(),ke(V,{key:0,class:"drawer-btn",onClick:p,quaternary:"",circle:"",size:"medium"},{icon:A(()=>[L(y,null,{default:A(()=>[L(X(ln))]),_:1})]),_:1})):re("",!0),e.back?(I(),ke(V,{key:1,class:"back-btn",onClick:w,quaternary:"",circle:"",size:"small"},{icon:A(()=>[L(y,null,{default:A(()=>[L(X(Qt))]),_:1})]),_:1})):re("",!0),ft(" "+bt(t.title)+" ",1),t.theme?(I(),ke(j,{key:2,value:X(n).state.theme==="dark","onUpdate:value":u,size:"small",class:"theme-switch-wrap"},{"checked-icon":A(()=>[L(y,{component:X(hn)},null,8,["component"])]),"unchecked-icon":A(()=>[L(y,{component:X(on)},null,8,["component"])]),_:1},8,["value"])):re("",!0)])]),_:1})],64)}}});export{mn as _};
