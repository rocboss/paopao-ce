import{cL as ne,cM as Be,cN as Se,bq as $e,r as D,k as ze,cO as Me,m as Re,c as ae,f as t,cB as oe,b as X,e as h,a as ie,d as L,u as Ve,x as le,o as Le,t as Fe,s as Te,y as E,z as x,br as Z,c7 as f,A as Oe,cP as G,h as r,B as y,cz as Ee,ce as Pe,w as J,W as z,Y as ee,Z as W,a2 as Ae,a6 as Q,a5 as R,a4 as P,a3 as A,a7 as re,a9 as Ne,aa as De,$ as He,ak as Ie,af as We,K as Ke,bP as Ue}from"./index-476ef983.js";let N=0;const je=typeof window<"u"&&window.matchMedia!==void 0,B=D(null);let c,C;function H(e){e.matches&&(B.value="dark")}function I(e){e.matches&&(B.value="light")}function qe(){c=window.matchMedia("(prefers-color-scheme: dark)"),C=window.matchMedia("(prefers-color-scheme: light)"),c.matches?B.value="dark":C.matches?B.value="light":B.value=null,c.addEventListener?(c.addEventListener("change",H),C.addEventListener("change",I)):c.addListener&&(c.addListener(H),C.addListener(I))}function Ye(){"removeEventListener"in c?(c.removeEventListener("change",H),C.removeEventListener("change",I)):"removeListener"in c&&(c.removeListener(H),C.removeListener(I)),c=void 0,C=void 0}let se=!0;function Xe(){return je?(N===0&&qe(),se&&(se=Be())&&(Se(()=>{N+=1}),$e(()=>{N-=1,N===0&&Ye()})),ne(B)):ne(B)}const Ze=e=>{const{primaryColor:o,opacityDisabled:i,borderRadius:s,textColor3:l}=e,b="rgba(0, 0, 0, .14)";return Object.assign(Object.assign({},Me),{iconColor:l,textColor:"white",loadingColor:o,opacityDisabled:i,railColor:b,railColorActive:o,buttonBoxShadow:"0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)",buttonColor:"#FFF",railBorderRadiusSmall:s,railBorderRadiusMedium:s,railBorderRadiusLarge:s,buttonBorderRadiusSmall:s,buttonBorderRadiusMedium:s,buttonBorderRadiusLarge:s,boxShadowFocus:`0 0 0 2px ${Re(o,{alpha:.2})}`})},Ge={name:"Switch",common:ze,self:Ze},Je=Ge,Qe=ae("switch",`
 height: var(--n-height);
 min-width: var(--n-width);
 vertical-align: middle;
 user-select: none;
 -webkit-user-select: none;
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
 `),ae("base-loading",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 font-size: calc(var(--n-button-width) - 4px);
 color: var(--n-loading-color);
 transition: color .3s var(--n-bezier);
 `,[oe({left:"50%",top:"50%",originalTransform:"translateX(-50%) translateY(-50%)"})]),t("checked, unchecked",`
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
 `),X("&:focus",[t("rail",`
 box-shadow: var(--n-box-shadow-focus);
 `)]),h("round",[t("rail","border-radius: calc(var(--n-rail-height) / 2);",[t("button","border-radius: calc(var(--n-button-height) / 2);")])]),ie("disabled",[ie("icon",[h("rubber-band",[h("pressed",[t("rail",[t("button","max-width: var(--n-button-width-pressed);")])]),t("rail",[X("&:active",[t("button","max-width: var(--n-button-width-pressed);")])]),h("active",[h("pressed",[t("rail",[t("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])]),t("rail",[X("&:active",[t("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])])])])])]),h("active",[t("rail",[t("button","left: calc(100% - var(--n-button-width) - var(--n-offset))")])]),t("rail",`
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
 `,[oe()]),t("button",`
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
 `)]),h("active",[t("rail","background-color: var(--n-rail-color-active);")]),h("loading",[t("rail",`
 cursor: wait;
 `)]),h("disabled",[t("rail",`
 cursor: not-allowed;
 opacity: .5;
 `)])]),et=Object.assign(Object.assign({},le.props),{size:{type:String,default:"medium"},value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,rubberBand:{type:Boolean,default:!0},onChange:[Function,Array]});let V;const tt=L({name:"Switch",props:et,setup(e){V===void 0&&(typeof CSS<"u"?typeof CSS.supports<"u"?V=CSS.supports("width","max(1px)"):V=!1:V=!0);const{mergedClsPrefixRef:o,inlineThemeDisabled:i}=Ve(e),s=le("Switch","-switch",Qe,Je,e,o),l=Le(e),{mergedSizeRef:b,mergedDisabledRef:d}=l,S=D(e.defaultValue),g=Fe(e,"value"),u=Te(g,S),k=E(()=>u.value===e.checkedValue),v=D(!1),a=D(!1),m=E(()=>{const{railStyle:n}=e;if(n)return n({focused:a.value,checked:k.value})});function w(n){const{"onUpdate:value":F,onChange:T,onUpdateValue:O}=e,{nTriggerFormInput:K,nTriggerFormChange:U}=l;F&&J(F,n),O&&J(O,n),T&&J(T,n),S.value=n,K(),U()}function ce(){const{nTriggerFormFocus:n}=l;n()}function de(){const{nTriggerFormBlur:n}=l;n()}function ue(){e.loading||d.value||(u.value!==e.checkedValue?w(e.checkedValue):w(e.uncheckedValue))}function he(){a.value=!0,ce()}function fe(){a.value=!1,de(),v.value=!1}function be(n){e.loading||d.value||n.key===" "&&(u.value!==e.checkedValue?w(e.checkedValue):w(e.uncheckedValue),v.value=!1)}function ge(n){e.loading||d.value||n.key===" "&&(n.preventDefault(),v.value=!0)}const te=E(()=>{const{value:n}=b,{self:{opacityDisabled:F,railColor:T,railColorActive:O,buttonBoxShadow:K,buttonColor:U,boxShadowFocus:ve,loadingColor:me,textColor:we,iconColor:pe,[x("buttonHeight",n)]:p,[x("buttonWidth",n)]:_e,[x("buttonWidthPressed",n)]:ke,[x("railHeight",n)]:_,[x("railWidth",n)]:M,[x("railBorderRadius",n)]:xe,[x("buttonBorderRadius",n)]:ye},common:{cubicBezierEaseInOut:Ce}}=s.value;let j,q,Y;return V?(j=`calc((${_} - ${p}) / 2)`,q=`max(${_}, ${p})`,Y=`max(${M}, calc(${M} + ${p} - ${_}))`):(j=Z((f(_)-f(p))/2),q=Z(Math.max(f(_),f(p))),Y=f(_)>f(p)?M:Z(f(M)+f(p)-f(_))),{"--n-bezier":Ce,"--n-button-border-radius":ye,"--n-button-box-shadow":K,"--n-button-color":U,"--n-button-width":_e,"--n-button-width-pressed":ke,"--n-button-height":p,"--n-height":q,"--n-offset":j,"--n-opacity-disabled":F,"--n-rail-border-radius":xe,"--n-rail-color":T,"--n-rail-color-active":O,"--n-rail-height":_,"--n-rail-width":M,"--n-width":Y,"--n-box-shadow-focus":ve,"--n-loading-color":me,"--n-text-color":we,"--n-icon-color":pe}}),$=i?Oe("switch",E(()=>b.value[0]),te,e):void 0;return{handleClick:ue,handleBlur:fe,handleFocus:he,handleKeyup:be,handleKeydown:ge,mergedRailStyle:m,pressed:v,mergedClsPrefix:o,mergedValue:u,checked:k,mergedDisabled:d,cssVars:i?void 0:te,themeClass:$==null?void 0:$.themeClass,onRender:$==null?void 0:$.onRender}},render(){const{mergedClsPrefix:e,mergedDisabled:o,checked:i,mergedRailStyle:s,onRender:l,$slots:b}=this;l==null||l();const{checked:d,unchecked:S,icon:g,"checked-icon":u,"unchecked-icon":k}=b,v=!(G(g)&&G(u)&&G(k));return r("div",{role:"switch","aria-checked":i,class:[`${e}-switch`,this.themeClass,v&&`${e}-switch--icon`,i&&`${e}-switch--active`,o&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`,this.rubberBand&&`${e}-switch--rubber-band`],tabindex:this.mergedDisabled?void 0:0,style:this.cssVars,onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},r("div",{class:`${e}-switch__rail`,"aria-hidden":"true",style:s},y(d,a=>y(S,m=>a||m?r("div",{"aria-hidden":!0,class:`${e}-switch__children-placeholder`},r("div",{class:`${e}-switch__rail-placeholder`},r("div",{class:`${e}-switch__button-placeholder`}),a),r("div",{class:`${e}-switch__rail-placeholder`},r("div",{class:`${e}-switch__button-placeholder`}),m)):null)),r("div",{class:`${e}-switch__button`},y(g,a=>y(u,m=>y(k,w=>r(Ee,null,{default:()=>this.loading?r(Pe,{key:"loading",clsPrefix:e,strokeWidth:20}):this.checked&&(m||a)?r("div",{class:`${e}-switch__button-icon`,key:m?"checked-icon":"icon"},m||a):!this.checked&&(w||a)?r("div",{class:`${e}-switch__button-icon`,key:w?"unchecked-icon":"icon"},w||a):null})))),y(d,a=>a&&r("div",{key:"checked",class:`${e}-switch__checked`},a)),y(S,a=>a&&r("div",{key:"unchecked",class:`${e}-switch__unchecked`},a)))))}}),nt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},at=W("path",{d:"M14.71 6.71a.996.996 0 0 0-1.41 0L8.71 11.3a.996.996 0 0 0 0 1.41l4.59 4.59a.996.996 0 1 0 1.41-1.41L10.83 12l3.88-3.88c.39-.39.38-1.03 0-1.41z",fill:"currentColor"},null,-1),ot=[at],it=L({name:"ChevronLeftRound",render:function(o,i){return z(),ee("svg",nt,ot)}}),rt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},st=W("path",{d:"M9.37 5.51A7.35 7.35 0 0 0 9.1 7.5c0 4.08 3.32 7.4 7.4 7.4c.68 0 1.35-.09 1.99-.27A7.014 7.014 0 0 1 12 19c-3.86 0-7-3.14-7-7c0-2.93 1.81-5.45 4.37-6.49zM12 3a9 9 0 1 0 9 9c0-.46-.04-.92-.1-1.36a5.389 5.389 0 0 1-4.4 2.26a5.403 5.403 0 0 1-3.14-9.8c-.44-.06-.9-.1-1.36-.1z",fill:"currentColor"},null,-1),lt=[st],ct=L({name:"DarkModeOutlined",render:function(o,i){return z(),ee("svg",rt,lt)}}),dt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 24 24"},ut=W("path",{d:"M12 9c1.65 0 3 1.35 3 3s-1.35 3-3 3s-3-1.35-3-3s1.35-3 3-3m0-2c-2.76 0-5 2.24-5 5s2.24 5 5 5s5-2.24 5-5s-2.24-5-5-5zM2 13h2c.55 0 1-.45 1-1s-.45-1-1-1H2c-.55 0-1 .45-1 1s.45 1 1 1zm18 0h2c.55 0 1-.45 1-1s-.45-1-1-1h-2c-.55 0-1 .45-1 1s.45 1 1 1zM11 2v2c0 .55.45 1 1 1s1-.45 1-1V2c0-.55-.45-1-1-1s-1 .45-1 1zm0 18v2c0 .55.45 1 1 1s1-.45 1-1v-2c0-.55-.45-1-1-1s-1 .45-1 1zM5.99 4.58a.996.996 0 0 0-1.41 0a.996.996 0 0 0 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0s.39-1.03 0-1.41L5.99 4.58zm12.37 12.37a.996.996 0 0 0-1.41 0a.996.996 0 0 0 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0a.996.996 0 0 0 0-1.41l-1.06-1.06zm1.06-10.96a.996.996 0 0 0 0-1.41a.996.996 0 0 0-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06zM7.05 18.36a.996.996 0 0 0 0-1.41a.996.996 0 0 0-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06z",fill:"currentColor"},null,-1),ht=[ut],ft=L({name:"LightModeOutlined",render:function(o,i){return z(),ee("svg",dt,ht)}}),bt={class:"navbar"},vt=L({__name:"main-nav",props:{title:{default:""},back:{type:Boolean,default:!1},theme:{type:Boolean,default:!0}},setup(e){const o=e,i=He(),s=Ie(),l=d=>{d?(localStorage.setItem("PAOPAO_THEME","dark"),i.commit("triggerTheme","dark")):(localStorage.setItem("PAOPAO_THEME","light"),i.commit("triggerTheme","light"))},b=()=>{window.history.length<=1?s.push({path:"/"}):s.go(-1)};return Ae(()=>{localStorage.getItem("PAOPAO_THEME")||l(Xe()==="dark")}),(d,S)=>{const g=We,u=Ke,k=tt,v=Ue;return z(),Q(v,{size:"small",bordered:!0,class:"nav-title-card"},{header:R(()=>[W("div",bt,[e.back?(z(),Q(u,{key:0,class:"back-btn",onClick:b,quaternary:"",circle:"",size:"small"},{icon:R(()=>[P(g,null,{default:R(()=>[P(A(it))]),_:1})]),_:1})):re("",!0),Ne(" "+De(o.title)+" ",1),o.theme?(z(),Q(k,{key:1,value:A(i).state.theme==="dark","onUpdate:value":l,size:"small",class:"theme-switch-wrap"},{"checked-icon":R(()=>[P(g,{component:A(ft)},null,8,["component"])]),"unchecked-icon":R(()=>[P(g,{component:A(ct)},null,8,["component"])]),_:1},8,["value"])):re("",!0)])]),_:1})}}});export{vt as _};
