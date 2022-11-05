import{ao as G,ap as Ce,aq as Be,ar as ze,r as N,c as Re,as as Se,b as Me,l as Q,m as t,at as ee,n as U,o as w,p as te,d as V,u as Ve,q as oe,f as $e,t as Le,g as Fe,s as T,v as _,au as q,av as h,y as Te,aw as K,h as r,z as k,ax as Oe,ay as Ee,k as Y,V as z,W as J,Y as W,Z as Ae,aj as Ne,a1 as Pe,a5 as X,a4 as M,a3 as O,a2 as E,a6 as ne,a8 as De,a9 as We,ae as Ie,J as He,az as je}from"./index.f480f018.js";let A=0;const Ue=typeof window!="undefined"&&window.matchMedia!==void 0,y=N(null);let c,x;function P(e){e.matches&&(y.value="dark")}function D(e){e.matches&&(y.value="light")}function qe(){c=window.matchMedia("(prefers-color-scheme: dark)"),x=window.matchMedia("(prefers-color-scheme: light)"),c.matches?y.value="dark":x.matches?y.value="light":y.value=null,c.addEventListener?(c.addEventListener("change",P),x.addEventListener("change",D)):c.addListener&&(c.addListener(P),x.addListener(D))}function Ke(){"removeEventListener"in c?(c.removeEventListener("change",P),x.removeEventListener("change",D)):"removeListener"in c&&(c.removeListener(P),x.removeListener(D)),c=void 0,x=void 0}let ae=!0;function Ye(){return Ue?(A===0&&qe(),ae&&(ae=Ce())&&(Be(()=>{A+=1}),ze(()=>{A-=1,A===0&&Ke()})),G(y)):G(y)}const Xe=e=>{const{primaryColor:o,opacityDisabled:i,borderRadius:s,textColor3:l}=e,f="rgba(0, 0, 0, .14)";return Object.assign(Object.assign({},Se),{iconColor:l,textColor:"white",loadingColor:o,opacityDisabled:i,railColor:f,railColorActive:o,buttonBoxShadow:"0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)",buttonColor:"#FFF",railBorderRadiusSmall:s,railBorderRadiusMedium:s,railBorderRadiusLarge:s,buttonBorderRadiusSmall:s,buttonBorderRadiusMedium:s,buttonBorderRadiusLarge:s,boxShadowFocus:`0 0 0 2px ${Me(o,{alpha:.2})}`})},Je={name:"Switch",common:Re,self:Xe};var Ze=Je,Ge=Q("switch",`
 height: var(--n-height);
 min-width: var(--n-width);
 vertical-align: middle;
 user-select: none;
 display: inline-flex;
 outline: none;
 justify-content: center;
 align-items: center;
`,[t("children-placeholder",`
 height: var(--n-rail-height);
 display: flex;
 flex-direction: column;
 overflow: hidden;
 pointer-events: none;
 visibility: hidden;
 `),t("rail-placeholder",`
 display: flex;
 flex-wrap: none;
 `),t("button-placeholder",`
 width: calc(1.75 * var(--n-rail-height));
 height: var(--n-rail-height);
 `),Q("base-loading",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 font-size: calc(var(--n-button-width) - 4px);
 color: var(--n-loading-color);
 transition: color .3s var(--n-bezier);
 `,[ee({originalTransform:"translateX(-50%) translateY(-50%)"})]),t("checked, unchecked",`
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
 `),t("checked",`
 right: 0;
 padding-right: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),t("unchecked",`
 left: 0;
 justify-content: flex-end;
 padding-left: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),U("&:focus",[t("rail",`
 box-shadow: var(--n-box-shadow-focus);
 `)]),w("round",[t("rail",{borderRadius:"calc(var(--n-rail-height) / 2)"},[t("button",{borderRadius:"calc(var(--n-button-height) / 2)"})])]),te("disabled",[te("icon",[w("pressed",[t("rail",[t("button",{maxWidth:"var(--n-button-width-pressed)"})])]),t("rail",[U("&:active",[t("button",{maxWidth:"var(--n-button-width-pressed)"})])]),w("active",[w("pressed",[t("rail",[t("button",{left:"calc(100% - var(--n-offset) - var(--n-button-width-pressed))"})])]),t("rail",[U("&:active",[t("button",{left:"calc(100% - var(--n-offset) - var(--n-button-width-pressed))"})])])])])]),w("active",[t("rail",[t("button",{left:"calc(100% - (var(--n-rail-height) + var(--n-button-width)) / 2)"})])]),t("rail",`
 overflow: hidden;
 height: var(--n-rail-height);
 min-width: var(--n-rail-width);
 border-radius: var(--n-rail-border-radius);
 cursor: pointer;
 position: relative;
 transition:
 background .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-rail-color);
 `,[t("button-icon",`
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
 `,[ee()]),t("button",`
 align-items: center; 
 top: var(--n-offset);
 left: var(--n-offset);
 height: var(--n-button-width);
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
 `)]),w("active",[t("rail","background-color: var(--n-rail-color-active);")]),w("disabled",[t("rail",`
 cursor: not-allowed;
 opacity: .5;
 `)]),w("loading",[t("rail",`
 pointer-events: none;
 `)])]);const Qe=Object.assign(Object.assign({},oe.props),{size:{type:String,default:"medium"},value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,onChange:[Function,Array]});var et=V({name:"Switch",props:Qe,setup(e){const{mergedClsPrefixRef:o,inlineThemeDisabled:i}=Ve(e),s=oe("Switch","-switch",Ge,Ze,e,o),l=$e(e),{mergedSizeRef:f,mergedDisabledRef:d}=l,C=N(e.defaultValue),v=Le(e,"value"),u=Fe(v,C),p=T(()=>u.value===e.checkedValue),g=N(!1),a=N(!1),m=T(()=>{const{railStyle:n}=e;if(!!n)return n({focused:a.value,checked:p.value})});function b(n){const{"onUpdate:value":$,onChange:L,onUpdateValue:F}=e,{nTriggerFormInput:I,nTriggerFormChange:H}=l;$&&Y($,n),F&&Y(F,n),L&&Y(L,n),C.value=n,I(),H()}function ie(){const{nTriggerFormFocus:n}=l;n()}function re(){const{nTriggerFormBlur:n}=l;n()}function se(){e.loading||d.value||(u.value!==e.checkedValue?b(e.checkedValue):b(e.uncheckedValue))}function le(){a.value=!0,ie()}function ce(){a.value=!1,re(),g.value=!1}function de(n){e.loading||d.value||n.code==="Space"&&(u.value!==e.checkedValue?b(e.checkedValue):b(e.uncheckedValue),g.value=!1)}function ue(n){e.loading||d.value||n.code==="Space"&&(n.preventDefault(),g.value=!0)}const Z=T(()=>{const{value:n}=f,{self:{opacityDisabled:$,railColor:L,railColorActive:F,buttonBoxShadow:I,buttonColor:H,boxShadowFocus:he,loadingColor:fe,textColor:ve,iconColor:ge,[_("buttonHeight",n)]:R,[_("buttonWidth",n)]:me,[_("buttonWidthPressed",n)]:be,[_("railHeight",n)]:S,[_("railWidth",n)]:j,[_("railBorderRadius",n)]:we,[_("buttonBorderRadius",n)]:pe},common:{cubicBezierEaseInOut:_e}}=s.value,ke=q((h(S)-h(R))/2),xe=q(Math.max(h(S),h(R))),ye=h(S)>h(R)?j:q(h(j)+h(R)-h(S));return{"--n-bezier":_e,"--n-button-border-radius":pe,"--n-button-box-shadow":I,"--n-button-color":H,"--n-button-width":me,"--n-button-width-pressed":be,"--n-button-height":R,"--n-height":xe,"--n-offset":ke,"--n-opacity-disabled":$,"--n-rail-border-radius":we,"--n-rail-color":L,"--n-rail-color-active":F,"--n-rail-height":S,"--n-rail-width":j,"--n-width":ye,"--n-box-shadow-focus":he,"--n-loading-color":fe,"--n-text-color":ve,"--n-icon-color":ge}}),B=i?Te("switch",T(()=>f.value[0]),Z,e):void 0;return{handleClick:se,handleBlur:ce,handleFocus:le,handleKeyup:de,handleKeydown:ue,mergedRailStyle:m,pressed:g,mergedClsPrefix:o,mergedValue:u,checked:p,mergedDisabled:d,cssVars:i?void 0:Z,themeClass:B==null?void 0:B.themeClass,onRender:B==null?void 0:B.onRender}},render(){const{mergedClsPrefix:e,mergedDisabled:o,checked:i,mergedRailStyle:s,onRender:l,$slots:f}=this;l==null||l();const{checked:d,unchecked:C,icon:v,"checked-icon":u,"unchecked-icon":p}=f,g=!(K(v)&&K(u)&&K(p));return r("div",{role:"switch","aria-checked":i,class:[`${e}-switch`,this.themeClass,g&&`${e}-switch--icon`,i&&`${e}-switch--active`,o&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`],tabindex:this.mergedDisabled?void 0:0,style:this.cssVars,onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},r("div",{class:`${e}-switch__rail`,"aria-hidden":"true",style:s},k(d,a=>k(C,m=>a||m?r("div",{"aria-hidden":!0,class:`${e}-switch__children-placeholder`},r("div",{class:`${e}-switch__rail-placeholder`},r("div",{class:`${e}-switch__button-placeholder`}),a),r("div",{class:`${e}-switch__rail-placeholder`},r("div",{class:`${e}-switch__button-placeholder`}),m)):null)),r("div",{class:`${e}-switch__button`},k(v,a=>k(u,m=>k(p,b=>r(Oe,null,{default:()=>this.loading?r(Ee,{key:"loading",clsPrefix:e,strokeWidth:20}):this.checked&&(m||a)?r("div",{class:`${e}-switch__button-icon`,key:m?"checked-icon":"icon"},m||a):!this.checked&&(b||a)?r("div",{class:`${e}-switch__button-icon`,key:b?"unchecked-icon":"icon"},b||a):null})))),k(d,a=>a&&r("div",{key:"checked",class:`${e}-switch__checked`},a)),k(C,a=>a&&r("div",{key:"unchecked",class:`${e}-switch__unchecked`},a)))))}});const tt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},nt=W("path",{d:"M14.71 6.71a.996.996 0 0 0-1.41 0L8.71 11.3a.996.996 0 0 0 0 1.41l4.59 4.59a.996.996 0 1 0 1.41-1.41L10.83 12l3.88-3.88c.39-.39.38-1.03 0-1.41z",fill:"currentColor"},null,-1),at=[nt];var ot=V({name:"ChevronLeftRound",render:function(o,i){return z(),J("svg",tt,at)}});const it={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},rt=W("path",{d:"M9.37 5.51A7.35 7.35 0 0 0 9.1 7.5c0 4.08 3.32 7.4 7.4 7.4c.68 0 1.35-.09 1.99-.27A7.014 7.014 0 0 1 12 19c-3.86 0-7-3.14-7-7c0-2.93 1.81-5.45 4.37-6.49zM12 3a9 9 0 1 0 9 9c0-.46-.04-.92-.1-1.36a5.389 5.389 0 0 1-4.4 2.26a5.403 5.403 0 0 1-3.14-9.8c-.44-.06-.9-.1-1.36-.1z",fill:"currentColor"},null,-1),st=[rt];var lt=V({name:"DarkModeOutlined",render:function(o,i){return z(),J("svg",it,st)}});const ct={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},dt=W("path",{d:"M12 9c1.65 0 3 1.35 3 3s-1.35 3-3 3s-3-1.35-3-3s1.35-3 3-3m0-2c-2.76 0-5 2.24-5 5s2.24 5 5 5s5-2.24 5-5s-2.24-5-5-5zM2 13h2c.55 0 1-.45 1-1s-.45-1-1-1H2c-.55 0-1 .45-1 1s.45 1 1 1zm18 0h2c.55 0 1-.45 1-1s-.45-1-1-1h-2c-.55 0-1 .45-1 1s.45 1 1 1zM11 2v2c0 .55.45 1 1 1s1-.45 1-1V2c0-.55-.45-1-1-1s-1 .45-1 1zm0 18v2c0 .55.45 1 1 1s1-.45 1-1v-2c0-.55-.45-1-1-1s-1 .45-1 1zM5.99 4.58a.996.996 0 0 0-1.41 0a.996.996 0 0 0 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0s.39-1.03 0-1.41L5.99 4.58zm12.37 12.37a.996.996 0 0 0-1.41 0a.996.996 0 0 0 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0a.996.996 0 0 0 0-1.41l-1.06-1.06zm1.06-10.96a.996.996 0 0 0 0-1.41a.996.996 0 0 0-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06zM7.05 18.36a.996.996 0 0 0 0-1.41a.996.996 0 0 0-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06z",fill:"currentColor"},null,-1),ut=[dt];var ht=V({name:"LightModeOutlined",render:function(o,i){return z(),J("svg",ct,ut)}});const ft={class:"navbar"},gt=V({props:{title:{default:""},back:{type:Boolean,default:!1},theme:{type:Boolean,default:!0}},setup(e){const o=e,i=Ae(),s=Ne(),l=d=>{d?(localStorage.setItem("PAOPAO_THEME","dark"),i.commit("triggerTheme","dark")):(localStorage.setItem("PAOPAO_THEME","light"),i.commit("triggerTheme","light"))},f=()=>{window.history.length<=1?s.push({path:"/"}):s.go(-1)};return Pe(()=>{localStorage.getItem("PAOPAO_THEME")||l(Ye()==="dark")}),(d,C)=>{const v=Ie,u=He,p=et,g=je;return z(),X(g,{size:"small",bordered:!0,class:"nav-title-card"},{header:M(()=>[W("div",ft,[e.back?(z(),X(u,{key:0,class:"back-btn",onClick:f,quaternary:"",circle:"",size:"small"},{icon:M(()=>[O(v,null,{default:M(()=>[O(E(ot))]),_:1})]),_:1})):ne("",!0),De(" "+We(o.title)+" ",1),o.theme?(z(),X(p,{key:1,value:E(i).state.theme==="dark","onUpdate:value":l,size:"small",class:"theme-switch-wrap"},{"checked-icon":M(()=>[O(v,{component:E(ht)},null,8,["component"])]),"unchecked-icon":M(()=>[O(v,{component:E(lt)},null,8,["component"])]),_:1},8,["value"])):ne("",!0)])]),_:1})}}});export{gt as _};
