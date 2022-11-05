import{c as j,ce as k,cf as u,b as f,l as y,m as c,bz as B,o as F,n as M,d as O,u as N,q as E,x as V,s as S,ba as K,v as g,y as q,r as D,h as s,b2 as G,cg as J,K as Q,N as U,bt as X,bu as Y,bv as Z,bs as ee,z as oe,bx as re}from"./index.f480f018.js";const ne=r=>{const{lineHeight:o,borderRadius:i,fontWeightStrong:m,baseColor:n,dividerColor:v,actionColor:z,textColor1:t,textColor2:a,closeColor:h,closeColorHover:C,closeColorPressed:b,infoColor:l,successColor:e,warningColor:p,errorColor:x,fontSize:T}=r;return Object.assign(Object.assign({},k),{fontSize:T,lineHeight:o,titleFontWeight:m,borderRadius:i,border:`1px solid ${v}`,color:z,titleTextColor:t,iconColor:a,contentTextColor:a,closeColor:h,closeColorHover:C,closeColorPressed:b,borderInfo:`1px solid ${u(n,f(l,{alpha:.25}))}`,colorInfo:u(n,f(l,{alpha:.08})),titleTextColorInfo:t,iconColorInfo:l,contentTextColorInfo:a,closeColorInfo:h,closeColorHoverInfo:C,closeColorPressedInfo:b,borderSuccess:`1px solid ${u(n,f(e,{alpha:.25}))}`,colorSuccess:u(n,f(e,{alpha:.08})),titleTextColorSuccess:t,iconColorSuccess:e,contentTextColorSuccess:a,closeColorSuccess:h,closeColorHoverSuccess:C,closeColorPressedSuccess:b,borderWarning:`1px solid ${u(n,f(p,{alpha:.33}))}`,colorWarning:u(n,f(p,{alpha:.08})),titleTextColorWarning:t,iconColorWarning:p,contentTextColorWarning:a,closeColorWarning:h,closeColorHoverWarning:C,closeColorPressedWarning:b,borderError:`1px solid ${u(n,f(x,{alpha:.25}))}`,colorError:u(n,f(x,{alpha:.08})),titleTextColorError:t,iconColorError:x,contentTextColorError:a,closeColorError:h,closeColorHoverError:C,closeColorPressedError:b})},te={name:"Alert",common:j,self:ne};var le=te,se=y("alert",`
 line-height: var(--n-line-height);
 border-radius: var(--n-border-radius);
 position: relative;
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-color);
 text-align: start;
`,[c("icon",{color:"var(--n-icon-color)"}),y("alert-body",{border:"var(--n-border)",padding:"var(--n-padding)"},[c("title",{color:"var(--n-title-text-color)"}),c("content",{color:"var(--n-content-text-color)"})]),B({originalTransition:"transform .3s var(--n-bezier)",enterToProps:{transform:"scale(1)"},leaveToProps:{transform:"scale(0.9)"}}),c("icon",`
 position: absolute;
 left: 0;
 top: 0;
 align-items: center;
 justify-content: center;
 display: flex;
 width: var(--n-icon-size);
 height: var(--n-icon-size);
 font-size: var(--n-icon-size);
 margin: var(--n-icon-margin);
 `),c("close",`
 transition: color .3s var(--n-bezier);
 position: absolute;
 right: 0;
 top: 0;
 margin: var(--n-close-margin);
 font-size: var(--n-close-size);
 `),F("show-icon",[y("alert-body",{paddingLeft:"calc(var(--n-icon-margin-left) + var(--n-icon-size) + var(--n-icon-margin-right))"})]),y("alert-body",`
 border-radius: var(--n-border-radius);
 transition: border-color .3s var(--n-bezier);
 `,[c("title",`
 transition: color .3s var(--n-bezier);
 font-size: 16px;
 line-height: 19px;
 font-weight: var(--n-title-font-weight);
 `,[M("& +",[c("content",{marginTop:"9px"})])]),c("content",{transition:"color .3s var(--n-bezier)",fontSize:"var(--n-font-size)"})]),c("icon",{transition:"color .3s var(--n-bezier)"})]);const ie=Object.assign(Object.assign({},E.props),{title:{type:String,default:void 0},showIcon:{type:Boolean,default:!0},type:{type:String,default:"default"},closable:{type:Boolean,default:!1},onClose:Function,onAfterLeave:{type:Function,default:void 0},onAfterHide:{type:Function,validator:()=>!0,default:void 0}});var ce=O({name:"Alert",inheritAttrs:!1,props:ie,setup(r){const{mergedClsPrefixRef:o,inlineThemeDisabled:i,mergedRtlRef:m}=N(r),n=E("Alert","-alert",se,le,r,o),v=V("Alert",m,o),z=S(()=>{const{common:{cubicBezierEaseInOut:l},self:e}=n.value,{fontSize:p,borderRadius:x,titleFontWeight:T,lineHeight:P,iconSize:$,iconMargin:I,iconMarginRtl:R,closeSize:_,closeMargin:A,closeMarginRtl:W,padding:w}=e,{type:d}=r,{left:H,right:L}=K(I);return{"--n-bezier":l,"--n-color":e[g("color",d)],"--n-close-color":e[g("closeColor",d)],"--n-close-color-hover":e[g("closeColorHover",d)],"--n-close-color-pressed":e[g("closeColorPressed",d)],"--n-icon-color":e[g("iconColor",d)],"--n-border":e[g("border",d)],"--n-title-text-color":e[g("titleTextColor",d)],"--n-content-text-color":e[g("contentTextColor",d)],"--n-line-height":P,"--n-border-radius":x,"--n-font-size":p,"--n-title-font-weight":T,"--n-icon-size":$,"--n-icon-margin":I,"--n-icon-margin-rtl":R,"--n-close-size":_,"--n-close-margin":A,"--n-close-margin-rtl":W,"--n-padding":w,"--n-icon-margin-left":H,"--n-icon-margin-right":L}}),t=i?q("alert",S(()=>r.type[0]),z,r):void 0,a=D(!0),h=()=>{const{onAfterLeave:l,onAfterHide:e}=r;l&&l(),e&&e()};return{rtlEnabled:v,mergedClsPrefix:o,visible:a,handleCloseClick:()=>{var l;Promise.resolve((l=r.onClose)===null||l===void 0?void 0:l.call(r)).then(e=>{e!==!1&&(a.value=!1)})},handleAfterLeave:()=>{h()},mergedTheme:n,cssVars:i?void 0:z,themeClass:t==null?void 0:t.themeClass,onRender:t==null?void 0:t.onRender}},render(){var r;return(r=this.onRender)===null||r===void 0||r.call(this),s(re,{onAfterLeave:this.handleAfterLeave},{default:()=>{const{mergedClsPrefix:o,$slots:i}=this,m={class:[`${o}-alert`,this.themeClass,this.showIcon&&`${o}-alert--show-icon`,this.rtlEnabled&&`${o}-alert--rtl`],style:this.cssVars,role:"alert"};return this.visible?s("div",Object.assign({},G(this.$attrs,m)),this.closable&&s(J,{clsPrefix:o,class:`${o}-alert__close`,onClick:this.handleCloseClick}),this.showIcon&&s("div",{class:`${o}-alert__icon`,"aria-hidden":"true"},Q(i.icon,()=>[s(U,{clsPrefix:o},{default:()=>{switch(this.type){case"success":return s(ee,null);case"info":return s(Z,null);case"warning":return s(Y,null);case"error":return s(X,null);default:return null}}})])),s("div",{class:`${o}-alert-body`},oe(i.header,n=>{const v=n||this.title;return v?s("div",{class:`${o}-alert-body__title`},v):null}),i.default&&s("div",{class:`${o}-alert-body__content`},i))):null}})}});export{ce as _};
