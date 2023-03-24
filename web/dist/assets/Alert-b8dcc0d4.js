import{k as M,cE as O,cF as u,m as f,c as P,f as i,e as E,cA as N,b as V,d as D,u as G,x as H,i as K,y as $,cd as q,z as a,A as J,r as Q,h as l,b7 as U,cG as X,L as Y,B as Z,cx as oo,N as eo,cu as ro,an as no,cv as so,ct as lo}from"./index-476ef983.js";const to=r=>{const{lineHeight:e,borderRadius:d,fontWeightStrong:b,baseColor:t,dividerColor:v,actionColor:S,textColor1:g,textColor2:s,closeColorHover:h,closeColorPressed:C,closeIconColor:m,closeIconColorHover:p,closeIconColorPressed:n,infoColor:o,successColor:I,warningColor:x,errorColor:z,fontSize:T}=r;return Object.assign(Object.assign({},O),{fontSize:T,lineHeight:e,titleFontWeight:b,borderRadius:d,border:`1px solid ${v}`,color:S,titleTextColor:g,iconColor:s,contentTextColor:s,closeBorderRadius:d,closeColorHover:h,closeColorPressed:C,closeIconColor:m,closeIconColorHover:p,closeIconColorPressed:n,borderInfo:`1px solid ${u(t,f(o,{alpha:.25}))}`,colorInfo:u(t,f(o,{alpha:.08})),titleTextColorInfo:g,iconColorInfo:o,contentTextColorInfo:s,closeColorHoverInfo:h,closeColorPressedInfo:C,closeIconColorInfo:m,closeIconColorHoverInfo:p,closeIconColorPressedInfo:n,borderSuccess:`1px solid ${u(t,f(I,{alpha:.25}))}`,colorSuccess:u(t,f(I,{alpha:.08})),titleTextColorSuccess:g,iconColorSuccess:I,contentTextColorSuccess:s,closeColorHoverSuccess:h,closeColorPressedSuccess:C,closeIconColorSuccess:m,closeIconColorHoverSuccess:p,closeIconColorPressedSuccess:n,borderWarning:`1px solid ${u(t,f(x,{alpha:.33}))}`,colorWarning:u(t,f(x,{alpha:.08})),titleTextColorWarning:g,iconColorWarning:x,contentTextColorWarning:s,closeColorHoverWarning:h,closeColorPressedWarning:C,closeIconColorWarning:m,closeIconColorHoverWarning:p,closeIconColorPressedWarning:n,borderError:`1px solid ${u(t,f(z,{alpha:.25}))}`,colorError:u(t,f(z,{alpha:.08})),titleTextColorError:g,iconColorError:z,contentTextColorError:s,closeColorHoverError:h,closeColorPressedError:C,closeIconColorError:m,closeIconColorHoverError:p,closeIconColorPressedError:n})},io={name:"Alert",common:M,self:to},co=io,ao=P("alert",`
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
 `),E("closable",[P("alert-body",[i("title",`
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
 `),E("show-icon",[P("alert-body",{paddingLeft:"calc(var(--n-icon-margin-left) + var(--n-icon-size) + var(--n-icon-margin-right))"})]),P("alert-body",`
 border-radius: var(--n-border-radius);
 transition: border-color .3s var(--n-bezier);
 `,[i("title",`
 transition: color .3s var(--n-bezier);
 font-size: 16px;
 line-height: 19px;
 font-weight: var(--n-title-font-weight);
 `,[V("& +",[i("content",{marginTop:"9px"})])]),i("content",{transition:"color .3s var(--n-bezier)",fontSize:"var(--n-font-size)"})]),i("icon",{transition:"color .3s var(--n-bezier)"})]),go=Object.assign(Object.assign({},H.props),{title:String,showIcon:{type:Boolean,default:!0},type:{type:String,default:"default"},bordered:{type:Boolean,default:!0},closable:Boolean,onClose:Function,onAfterLeave:Function,onAfterHide:Function}),uo=D({name:"Alert",inheritAttrs:!1,props:go,setup(r){const{mergedClsPrefixRef:e,mergedBorderedRef:d,inlineThemeDisabled:b,mergedRtlRef:t}=G(r),v=H("Alert","-alert",ao,co,r,e),S=K("Alert",t,e),g=$(()=>{const{common:{cubicBezierEaseInOut:n},self:o}=v.value,{fontSize:I,borderRadius:x,titleFontWeight:z,lineHeight:T,iconSize:R,iconMargin:y,iconMarginRtl:_,closeIconSize:A,closeBorderRadius:W,closeSize:w,closeMargin:B,closeMarginRtl:L,padding:k}=o,{type:c}=r,{left:F,right:j}=q(y);return{"--n-bezier":n,"--n-color":o[a("color",c)],"--n-close-icon-size":A,"--n-close-border-radius":W,"--n-close-color-hover":o[a("closeColorHover",c)],"--n-close-color-pressed":o[a("closeColorPressed",c)],"--n-close-icon-color":o[a("closeIconColor",c)],"--n-close-icon-color-hover":o[a("closeIconColorHover",c)],"--n-close-icon-color-pressed":o[a("closeIconColorPressed",c)],"--n-icon-color":o[a("iconColor",c)],"--n-border":o[a("border",c)],"--n-title-text-color":o[a("titleTextColor",c)],"--n-content-text-color":o[a("contentTextColor",c)],"--n-line-height":T,"--n-border-radius":x,"--n-font-size":I,"--n-title-font-weight":z,"--n-icon-size":R,"--n-icon-margin":y,"--n-icon-margin-rtl":_,"--n-close-size":w,"--n-close-margin":B,"--n-close-margin-rtl":L,"--n-padding":k,"--n-icon-margin-left":F,"--n-icon-margin-right":j}}),s=b?J("alert",$(()=>r.type[0]),g,r):void 0,h=Q(!0),C=()=>{const{onAfterLeave:n,onAfterHide:o}=r;n&&n(),o&&o()};return{rtlEnabled:S,mergedClsPrefix:e,mergedBordered:d,visible:h,handleCloseClick:()=>{var n;Promise.resolve((n=r.onClose)===null||n===void 0?void 0:n.call(r)).then(o=>{o!==!1&&(h.value=!1)})},handleAfterLeave:()=>{C()},mergedTheme:v,cssVars:b?void 0:g,themeClass:s==null?void 0:s.themeClass,onRender:s==null?void 0:s.onRender}},render(){var r;return(r=this.onRender)===null||r===void 0||r.call(this),l(oo,{onAfterLeave:this.handleAfterLeave},{default:()=>{const{mergedClsPrefix:e,$slots:d}=this,b={class:[`${e}-alert`,this.themeClass,this.closable&&`${e}-alert--closable`,this.showIcon&&`${e}-alert--show-icon`,this.rtlEnabled&&`${e}-alert--rtl`],style:this.cssVars,role:"alert"};return this.visible?l("div",Object.assign({},U(this.$attrs,b)),this.closable&&l(X,{clsPrefix:e,class:`${e}-alert__close`,onClick:this.handleCloseClick}),this.bordered&&l("div",{class:`${e}-alert__border`}),this.showIcon&&l("div",{class:`${e}-alert__icon`,"aria-hidden":"true"},Y(d.icon,()=>[l(eo,{clsPrefix:e},{default:()=>{switch(this.type){case"success":return l(lo,null);case"info":return l(so,null);case"warning":return l(no,null);case"error":return l(ro,null);default:return null}}})])),l("div",{class:[`${e}-alert-body`,this.mergedBordered&&`${e}-alert-body--bordered`]},Z(d.header,t=>{const v=t||this.title;return v?l("div",{class:`${e}-alert-body__title`},v):null}),d.default&&l("div",{class:`${e}-alert-body__content`},d))):null}})}});export{uo as _};
