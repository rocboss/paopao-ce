import{k as M,cm as O,cn as u,m as v,c as P,f as i,e as H,bM as N,b as V,d as K,u as D,x as E,j as G,y as $,bm as q,z as c,A as J,r as Q,h as s,bd as U,co as X,L as Y,B as Z,bK as oo,N as eo,bG as ro,bH as no,bI as lo,bF as so}from"./index.d4f5aad2.js";const to=r=>{const{lineHeight:e,borderRadius:d,fontWeightStrong:C,baseColor:t,dividerColor:b,actionColor:S,textColor1:g,textColor2:l,closeColorHover:h,closeColorPressed:f,closeIconColor:m,closeIconColorHover:p,closeIconColorPressed:n,infoColor:o,successColor:I,warningColor:x,errorColor:z,fontSize:T}=r;return Object.assign(Object.assign({},O),{fontSize:T,lineHeight:e,titleFontWeight:C,borderRadius:d,border:`1px solid ${b}`,color:S,titleTextColor:g,iconColor:l,contentTextColor:l,closeBorderRadius:d,closeColorHover:h,closeColorPressed:f,closeIconColor:m,closeIconColorHover:p,closeIconColorPressed:n,borderInfo:`1px solid ${u(t,v(o,{alpha:.25}))}`,colorInfo:u(t,v(o,{alpha:.08})),titleTextColorInfo:g,iconColorInfo:o,contentTextColorInfo:l,closeColorHoverInfo:h,closeColorPressedInfo:f,closeIconColorInfo:m,closeIconColorHoverInfo:p,closeIconColorPressedInfo:n,borderSuccess:`1px solid ${u(t,v(I,{alpha:.25}))}`,colorSuccess:u(t,v(I,{alpha:.08})),titleTextColorSuccess:g,iconColorSuccess:I,contentTextColorSuccess:l,closeColorHoverSuccess:h,closeColorPressedSuccess:f,closeIconColorSuccess:m,closeIconColorHoverSuccess:p,closeIconColorPressedSuccess:n,borderWarning:`1px solid ${u(t,v(x,{alpha:.33}))}`,colorWarning:u(t,v(x,{alpha:.08})),titleTextColorWarning:g,iconColorWarning:x,contentTextColorWarning:l,closeColorHoverWarning:h,closeColorPressedWarning:f,closeIconColorWarning:m,closeIconColorHoverWarning:p,closeIconColorPressedWarning:n,borderError:`1px solid ${u(t,v(z,{alpha:.25}))}`,colorError:u(t,v(z,{alpha:.08})),titleTextColorError:g,iconColorError:z,contentTextColorError:l,closeColorHoverError:h,closeColorPressedError:f,closeIconColorError:m,closeIconColorHoverError:p,closeIconColorPressedError:n})},io={name:"Alert",common:M,self:to};var ao=io,co=P("alert",`
 line-height: var(--n-line-height);
 border-radius: var(--n-border-radius);
 position: relative;
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-color);
 text-align: start;
 word-break: break-word;
`,[i("border",`
 border-radius: inherit;
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 transition: border-color .3s var(--n-bezier);
 border: var(--n-border);
 pointer-events: none;
 `),H("closable",[P("alert-body",[i("title",`
 padding-right: 24px;
 `)])]),i("icon",{color:"var(--n-icon-color)"}),P("alert-body",{padding:"var(--n-padding)"},[i("title",{color:"var(--n-title-text-color)"}),i("content",{color:"var(--n-content-text-color)"})]),N({originalTransition:"transform .3s var(--n-bezier)",enterToProps:{transform:"scale(1)"},leaveToProps:{transform:"scale(0.9)"}}),i("icon",`
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
 `),i("close",`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 position: absolute;
 right: 0;
 top: 0;
 margin: var(--n-close-margin);
 `),H("show-icon",[P("alert-body",{paddingLeft:"calc(var(--n-icon-margin-left) + var(--n-icon-size) + var(--n-icon-margin-right))"})]),P("alert-body",`
 border-radius: var(--n-border-radius);
 transition: border-color .3s var(--n-bezier);
 `,[i("title",`
 transition: color .3s var(--n-bezier);
 font-size: 16px;
 line-height: 19px;
 font-weight: var(--n-title-font-weight);
 `,[V("& +",[i("content",{marginTop:"9px"})])]),i("content",{transition:"color .3s var(--n-bezier)",fontSize:"var(--n-font-size)"})]),i("icon",{transition:"color .3s var(--n-bezier)"})]);const go=Object.assign(Object.assign({},E.props),{title:String,showIcon:{type:Boolean,default:!0},type:{type:String,default:"default"},bordered:{type:Boolean,default:!0},closable:Boolean,onClose:Function,onAfterLeave:Function,onAfterHide:Function});var uo=K({name:"Alert",inheritAttrs:!1,props:go,setup(r){const{mergedClsPrefixRef:e,mergedBorderedRef:d,inlineThemeDisabled:C,mergedRtlRef:t}=D(r),b=E("Alert","-alert",co,ao,r,e),S=G("Alert",t,e),g=$(()=>{const{common:{cubicBezierEaseInOut:n},self:o}=b.value,{fontSize:I,borderRadius:x,titleFontWeight:z,lineHeight:T,iconSize:R,iconMargin:y,iconMarginRtl:_,closeIconSize:A,closeBorderRadius:W,closeSize:w,closeMargin:B,closeMarginRtl:L,padding:k}=o,{type:a}=r,{left:j,right:F}=q(y);return{"--n-bezier":n,"--n-color":o[c("color",a)],"--n-close-icon-size":A,"--n-close-border-radius":W,"--n-close-color-hover":o[c("closeColorHover",a)],"--n-close-color-pressed":o[c("closeColorPressed",a)],"--n-close-icon-color":o[c("closeIconColor",a)],"--n-close-icon-color-hover":o[c("closeIconColorHover",a)],"--n-close-icon-color-pressed":o[c("closeIconColorPressed",a)],"--n-icon-color":o[c("iconColor",a)],"--n-border":o[c("border",a)],"--n-title-text-color":o[c("titleTextColor",a)],"--n-content-text-color":o[c("contentTextColor",a)],"--n-line-height":T,"--n-border-radius":x,"--n-font-size":I,"--n-title-font-weight":z,"--n-icon-size":R,"--n-icon-margin":y,"--n-icon-margin-rtl":_,"--n-close-size":w,"--n-close-margin":B,"--n-close-margin-rtl":L,"--n-padding":k,"--n-icon-margin-left":j,"--n-icon-margin-right":F}}),l=C?J("alert",$(()=>r.type[0]),g,r):void 0,h=Q(!0),f=()=>{const{onAfterLeave:n,onAfterHide:o}=r;n&&n(),o&&o()};return{rtlEnabled:S,mergedClsPrefix:e,mergedBordered:d,visible:h,handleCloseClick:()=>{var n;Promise.resolve((n=r.onClose)===null||n===void 0?void 0:n.call(r)).then(o=>{o!==!1&&(h.value=!1)})},handleAfterLeave:()=>{f()},mergedTheme:b,cssVars:C?void 0:g,themeClass:l==null?void 0:l.themeClass,onRender:l==null?void 0:l.onRender}},render(){var r;return(r=this.onRender)===null||r===void 0||r.call(this),s(oo,{onAfterLeave:this.handleAfterLeave},{default:()=>{const{mergedClsPrefix:e,$slots:d}=this,C={class:[`${e}-alert`,this.themeClass,this.closable&&`${e}-alert--closable`,this.showIcon&&`${e}-alert--show-icon`,this.rtlEnabled&&`${e}-alert--rtl`],style:this.cssVars,role:"alert"};return this.visible?s("div",Object.assign({},U(this.$attrs,C)),this.closable&&s(X,{clsPrefix:e,class:`${e}-alert__close`,onClick:this.handleCloseClick}),this.bordered&&s("div",{class:`${e}-alert__border`}),this.showIcon&&s("div",{class:`${e}-alert__icon`,"aria-hidden":"true"},Y(d.icon,()=>[s(eo,{clsPrefix:e},{default:()=>{switch(this.type){case"success":return s(so,null);case"info":return s(lo,null);case"warning":return s(no,null);case"error":return s(ro,null);default:return null}}})])),s("div",{class:[`${e}-alert-body`,this.mergedBordered&&`${e}-alert-body--bordered`]},Z(d.header,t=>{const b=t||this.title;return b?s("div",{class:`${e}-alert-body__title`},b):null}),d.default&&s("div",{class:`${e}-alert-body__content`},d))):null}})}});export{uo as _};
