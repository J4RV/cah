(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{230:function(e,t,a){e.exports=a(397)},231:function(e,t,a){},397:function(e,t,a){"use strict";a.r(t);var n=Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));function r(e,t){navigator.serviceWorker.register(e).then(function(e){e.onupdatefound=function(){var a=e.installing;null!=a&&(a.onstatechange=function(){"installed"===a.state&&(navigator.serviceWorker.controller?(console.log("New content is available and will be used when all tabs for this page are closed. See http://bit.ly/CRA-PWA."),t&&t.onUpdate&&t.onUpdate(e)):(console.log("Content is cached for offline use."),t&&t.onSuccess&&t.onSuccess(e)))})}}).catch(function(e){console.error("Error during service worker registration:",e)})}var i=a(17),o=a(18),c=a(20),s=a(19),l=a(21),u=(a(231),a(7)),m=a(0),p=a.n(m),d=a(402),h=a(399),g=a(401),f=a(400),b=a(80),E=a.n(b),v=a(34),y=a.n(v);var O=function(e){return p.a.createElement(y.a,e,p.a.createElement("path",{d:"M12.007 0C6.12 0 1.1 4.27.157 10.08c-.944 5.813 2.468 11.45 8.054 13.312.19.064.397.033.555-.084.16-.117.25-.304.244-.5v-2.042c-3.33.735-4.037-1.56-4.037-1.56-.22-.726-.694-1.35-1.334-1.756-1.096-.75.074-.735.074-.735.773.103 1.454.557 1.846 1.23.694 1.21 2.23 1.638 3.45.96.056-.61.327-1.178.766-1.605-2.67-.3-5.462-1.335-5.462-6.002-.02-1.193.42-2.35 1.23-3.226-.327-1.015-.27-2.116.166-3.09 0 0 1.006-.33 3.3 1.23 1.966-.538 4.04-.538 6.003 0 2.295-1.5 3.3-1.23 3.3-1.23.445 1.006.49 2.144.12 3.18.81.877 1.25 2.033 1.23 3.226 0 4.607-2.805 5.627-5.476 5.927.578.583.88 1.386.825 2.206v3.29c-.005.2.092.393.26.507.164.115.377.14.565.063 5.568-1.88 8.956-7.514 8.007-13.313C22.892 4.267 17.884.007 12.008 0z"}))},j=a(25),w=a.n(j),C=a(44),k=a.n(C),S=a(398),x=a(79),N=a.n(x),R=a(164),I=a.n(R),P=a(41),A=a.n(P),W=function(e){return"gamestate/".concat(e,"/play-cards")},B="game/create",z="game/join",L=function(e){return"game/".concat(e,"/room-state")},T="game/start",D=function(e){var t=e.onElementClick;return p.a.createElement(p.a.Fragment,null,p.a.createElement(S.a,{to:"/game/list/my-games-in-progress"},p.a.createElement(A.a,{onClick:t},"My games in progress")),p.a.createElement(S.a,{to:"/game/list/open"},p.a.createElement(A.a,{onClick:t},"Open games")),p.a.createElement("a",{href:"user/logout"},p.a.createElement(A.a,{onClick:t},"Logout")))},G=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),o=0;o<n;o++)r[o]=arguments[o];return(a=Object(c.a)(this,(e=Object(s.a)(t)).call.apply(e,[this].concat(r)))).state={open:!1},a.handleToggle=function(){a.setState(function(e){return{open:!e.open}})},a.handleClose=function(e){a.anchorEl.contains(e.target)||a.setState({open:!1})},a}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this,t=this.props,a=t.classes,n=t.width,r=this.state.open;return Object(j.isWidthUp)("md",n,!0)?p.a.createElement(D,null):p.a.createElement("div",null,p.a.createElement(k.a,{className:a.menuButton,color:"inherit","aria-label":"Menu",onClick:this.handleToggle,buttonRef:function(t){return e.anchorEl=t}},p.a.createElement(I.a,null)),p.a.createElement(N.a,{open:r,anchorEl:this.anchorEl,onClose:this.handleClose},p.a.createElement(D,{onElementClick:this.handleClose})))}}]),t}(p.a.Component),M=w()()(Object(u.withStyles)(function(e){return{menuButton:{marginRight:-12,marginLeft:2*e.spacing.unit}}})(G)),F=a(94),H=a.n(F),U=a(81),_=a.n(U),q=a(10),J=a.n(q),V=a(22);var Y=Object(V.b)(function(e){return{username:e.username}})(w()()(Object(u.withStyles)(function(e){return{appbar:{color:e.palette.blackcard.text,backgroundColor:e.palette.blackcard.background},title:{flexGrow:1},user:{margin:e.spacing.unit},icon:{margin:e.spacing.unit,color:e.palette.grey[50]}}})(function(e){var t=e.username,a=e.title,n=e.shortTitle,r=e.width,i=e.classes;return p.a.createElement("div",null,p.a.createElement(E.a,{position:"static",className:i.appbar},p.a.createElement(_.a,null,p.a.createElement(J.a,{variant:"h6",color:"inherit",className:i.user},t),p.a.createElement(J.a,{variant:"h6",color:"inherit",className:i.title},"xs"===r?n:a),p.a.createElement(J.a,null,p.a.createElement("a",{target:"blank",href:"https://github.com/J4RV"},p.a.createElement(O,{className:i.icon}))),p.a.createElement(J.a,null,p.a.createElement("a",{target:"blank",href:"https://store.cardsagainsthumanity.com"},p.a.createElement(H.a,{className:i.icon}))),p.a.createElement(M,null))))}))),X=a(91),$=a.n(X),K=a(23),Q=a(29),Z=a.n(Q),ee=a(82),te=a.n(ee),ae=a(45),ne=a.n(ae),re=a(26),ie=a.n(re),oe=function(e){return{type:"PUSH_ERROR",payload:{msg:null!=e.response?e.response.data:e.toString()}}},ce=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),o=0;o<n;o++)r[o]=arguments[o];return(a=Object(c.a)(this,(e=Object(s.a)(t)).call.apply(e,[this].concat(r)))).state={name:"",password:"",waitingResponse:!1,createdSuccessfully:!1},a.handleSubmit=function(e){e.preventDefault(),a.setState(Object(K.a)({},a.state,{waitingResponse:!0}));var t={name:a.state.name,password:a.state.password};return ie.a.post(B,t).then(a.createdSuccessfully).catch(function(e){a.props.pushError(e),a.setState(Object(K.a)({},a.state,{waitingResponse:!1}))}),!1},a.createdSuccessfully=function(){a.setState(Object(K.a)({},a.state,{createdSuccessfully:!0}))},a.handleChangeName=function(e){var t=Object.assign({},a.state);t.name=e.target.value.trim(),a.setState(t)},a.handleChangePass=function(e){var t=Object.assign({},a.state);t.password=e.target.value.trim(),a.setState(t)},a}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this.props,t=e.classes,a=e.history;return this.state.createdSuccessfully?p.a.createElement(g.a,{to:"/game/list/open"}):p.a.createElement("form",{className:t.form,onSubmit:this.handleSubmit},p.a.createElement(J.a,{variant:"h6",className:t.formLabel},"Create a new game"),p.a.createElement(ne.a,{required:!0,fullWidth:!0,margin:"normal",label:"Room name",autoComplete:"roomName",s:!0,onChange:this.handleChangeName}),p.a.createElement(te.a,null,p.a.createElement(Z.a,{margin:"normal",onClick:a.goBack,disabled:this.state.waitingResponse},"Cancel"),p.a.createElement(Z.a,{margin:"normal",autoFocus:!0,type:"submit",variant:"contained",color:"primary",disabled:this.state.waitingResponse},"Create Game")))}}]),t}(m.Component),se=Object(V.b)(null,{pushError:oe})(Object(u.withStyles)(function(e){return{form:{padding:2*e.spacing.unit,maxWidth:360,marginLeft:"auto",marginRight:"auto"}}})(ce)),le=function(e){var t=e.className;return p.a.createElement(S.a,{to:"/game/list/open"},p.a.createElement(Z.a,{className:t},"Back to games list"))},ue=a(40),me=a.n(ue),pe=a(170),de=a(33),he=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),o=0;o<n;o++)r[o]=arguments[o];return(a=Object(c.a)(this,(e=Object(s.a)(t)).call.apply(e,[this].concat(r)))).randomRotate=function(){a.setState({rotation:5*Math.random()-2.5})},a}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e,t=this.props,a=t.text,n=t.isBlack,r=t.elevated,i=t.glowing,o=t.inHand,c=t.expansion,s=t.className,l=t.classes,u=t.style,m=Object(pe.a)(t,["text","isBlack","elevated","glowing","inHand","expansion","className","classes","style"]);e=i?l.glowing:r?l.floating:l.inTable;var d=n?l.black:l.white;return p.a.createElement("div",Object.assign({style:Object(K.a)({transform:"rotate(".concat(this.state.rotation,"deg)")},u),className:"".concat(l.card," ").concat(l.text," ").concat(d," ").concat(e,"    \n        ").concat(o?l.inHand:"","\n        ").concat(s||"")},m),p.a.createElement("div",null,a),p.a.createElement("div",{className:l.expansion},c))}},{key:"componentWillMount",value:function(){this.randomRotate()}}]),t}(p.a.Component),ge=w()()(Object(u.withStyles)(function(e){return{card:Object(de.a)({position:"relative",display:"inline-block",padding:e.spacing.unit,width:"8rem",height:"10rem",borderRadius:10,textAlign:"center",verticalAlign:"top",transformOrigin:"50% 80%"},e.breakpoints.down("sm"),{padding:.5*e.spacing.unit,width:"6.4rem",height:"8rem",borderRadius:8}),inHand:{margin:"0 0 -8px 0"},text:Object(de.a)({fontFamily:'"Open Sans", "Roboto", "Helvetica", "Arial", sans-serif',fontWeight:"600",fontSize:".8rem",whiteSpace:"pre-wrap"},e.breakpoints.down("sm"),{fontSize:".64rem"}),black:{color:e.palette.blackcard.text,background:e.palette.blackcard.background},white:{color:e.palette.whitecard.text,background:e.palette.whitecard.background},expansion:Object(de.a)({position:"absolute",bottom:e.spacing.unit,right:2*e.spacing.unit,marginLeft:e.spacing.unit,color:e.palette.expansion,fontSize:".8em",textAlign:"right"},e.breakpoints.down("sm"),{right:e.spacing.unit}),inTable:{boxShadow:e.shadows[1]},floating:{boxShadow:e.shadows[10]},glowing:{boxShadow:e.lights.glow}}})(he)),fe=a(165),be=a.n(fe),Ee=a(83),ve=a.n(Ee),ye=function(e){var t=e.classes,a=e.width,n=e.playCards;return"sm"===a||"xs"===a?p.a.createElement(ve.a,{"aria-label":"Play selected cards",color:"primary",onClick:n,className:t.smallScreenButton},p.a.createElement(be.a,null)):p.a.createElement(Z.a,{variant:"contained",color:"primary",onClick:n,className:t.largeScreenButton},"Play cards")};ye=w()()(ye);var Oe=function(e){var t=e.state;if(t.myPlayer.id===t.currentCzarID)return null;var a=t.blackCardInPlay.blanks-t.myPlayer.whiteCardsInPlay.length;return 0===a?null:p.a.createElement(J.a,{variant:"h6",gutterBottom:!0},"Play ",a," cards")},je=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),o=0;o<n;o++)r[o]=arguments[o];return(a=Object(c.a)(this,(e=Object(s.a)(t)).call.apply(e,[this].concat(r)))).state={cardIndexes:[],errormsg:null},a.canPlayCards=function(){var e=a.props.gamestate,t=e.currentCzarID===e.myPlayer.id,n="Sinners playing their cards"===e.phase;return!t&&n},a.handleCardClick=function(e){if(a.canPlayCards()){var t=a.state.cardIndexes.slice();t.includes(e)?t.splice(t.indexOf(e),1):t.push(e),a.setState({cardIndexes:t})}},a.playCards=function(){var e=a.props.gamestate;ie.a.post(W(e.id),{cardIndexes:a.state.cardIndexes}).then(function(e){a.setState({cardIndexes:[]})}).catch(function(e){return a.props.pushError(e)})},a}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this,t=this.props,a=t.gamestate,n=t.classes;return p.a.createElement("div",{className:n.hand},p.a.createElement(Oe,{state:a}),p.a.createElement("div",{className:n.cardsInHand},a.myPlayer.hand.map(function(t,a){return p.a.createElement(ge,Object.assign({},t,{handIndex:a,elevated:!0,inHand:!0,glowing:e.state.cardIndexes.includes(a),onClick:function(){return e.handleCardClick(a)}}))})),this.canPlayCards()?p.a.createElement(ye,{playCards:this.playCards,classes:n}):null)}}]),t}(m.Component),we=Object(V.b)(function(){},{pushError:oe})(w()()(Object(u.withStyles)(function(e){return{hand:{maxWidth:800,boxShadow:e.shadows[8],background:"#0004",padding:e.spacing.unit,paddingTop:2*e.spacing.unit,marginLeft:"auto",marginRight:"auto",textAlign:"center"},cardsInHand:{textAlign:"center",paddingBottom:8},largeScreenButton:{marginTop:e.spacing.unit},smallScreenButton:{position:"fixed",right:e.spacing.unit,bottom:e.spacing.unit}}})(je))),Ce=function(e){var t=e.player,a=e.itsYou,n=e.isCzar,r=e.classes;return p.a.createElement("div",{className:r.playerInfo},p.a.createElement("div",null,t.name," ",a?p.a.createElement("b",null,"(You)"):null),p.a.createElement("div",null,t.points.length," points"),p.a.createElement("div",null,n?p.a.createElement("b",null,"Current Czar"):"".concat(t.whiteCardsInPlay," card(s) in play")))},ke=w()()(Object(u.withStyles)(function(e){return{container:Object(de.a)({display:"flex",flexWrap:"wrap"},e.breakpoints.up("md"),{position:"fixed",right:8,top:72}),playerInfo:{background:"#EEE8",color:"#111",margin:4,padding:4,borderRadius:3,boxShadow:e.shadows[8],flexGrow:1}}})(function(e){var t=e.state,a=e.classes;return p.a.createElement(J.a,null,p.a.createElement("div",{className:a.container},t.players.map(function(e){return p.a.createElement(Ce,{player:e,itsYou:e.id===t.myPlayer.id,isCzar:e.id===t.currentCzarID,classes:a})})))})),Se=void 0,xe=function(e){var t=e.stateID,a=e.play,n=e.isCzar,r=e.classes,i=a.whiteCards;return null==i||0===i.length?null:p.a.createElement("div",{className:r.playerPlay},i.map(function(e){return p.a.createElement(ge,Object.assign({},e,{onClick:function(){return n&&function(e,t){ie.a.post(function(e){return"gamestate/".concat(e,"/choose-winner")}(e),{winner:t}).catch(function(e){return Se.props.pushError(e)})}(t,a.id)}}))}))},Ne=Object(V.b)(function(){},{pushError:oe})(Object(u.withStyles)(function(e){return{playerPlay:{marginLeft:2*e.spacing.unit,marginBottom:2*e.spacing.unit,display:"inline-block"}}})(function(e){var t=e.state,a=e.classes,n=t.myPlayer.id===t.currentCzarID;return t.sinnerPlays.length>0?p.a.createElement(p.a.Fragment,null,p.a.createElement(J.a,{variant:"h6",gutterBottom:!0},n?"Choose the winner":"Czar choosing winner..."),t.sinnerPlays.map(function(e){return p.a.createElement(xe,{stateID:t.id,play:e,isCzar:n,classes:a})})):p.a.createElement(J.a,{variant:"h6",gutterBottom:!0},n?p.a.createElement("p",null,"You are the Czar!"):null,"Waiting for all players to play their cards...")})),Re=function(e){var t=e.state,a=t.blackCardInPlay;return null==a?null:p.a.createElement("div",{className:"cah-table"},p.a.createElement(ge,Object.assign({},a,{isBlack:!0,style:{margin:"0 1rem 1rem 1rem"}})),p.a.createElement(Ne,{state:t}))},Ie=function(e){function t(){return Object(i.a)(this,t),Object(c.a)(this,Object(s.a)(t).apply(this,arguments))}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){return null==this.state?null:p.a.createElement("div",{className:"cah-game"},p.a.createElement(ke,{state:this.state}),p.a.createElement(Re,{state:this.state}),p.a.createElement(we,{gamestate:this.state}))}},{key:"componentWillMount",value:function(){var e=this,t=this.props.stateID;new WebSocket(function(e){return("http:"===document.location.protocol?"ws:":"wss:")+"//".concat(window.location.host,"/rest/gamestate/").concat(e,"/state-websocket")}(t)).onmessage=function(t){console.debug("updating game state",t.data),e.setState(JSON.parse(t.data))}}}]),t}(m.Component),Pe=Object(V.b)(null,{pushError:oe})(Ie),Ae=a(86),We=a.n(Ae),Be=a(84),ze=a.n(Be),Le=a(32),Te=a.n(Le),De=a(61),Ge=a.n(De),Me=a(62),Fe=a.n(Me),He={PaperProps:{style:{maxHeight:224,width:250}}},Ue=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),o=0;o<n;o++)r[o]=arguments[o];return(a=Object(c.a)(this,(e=Object(s.a)(t)).call.apply(e,[this].concat(r)))).state={expansions:["Loading..."],selected:[]},a.handleChangeSelect=function(e){var t=e.target.value;a.setState(Object(K.a)({},a.state,{selected:t})),a.props.onSelectedChange(t)},a}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this.props.classes,t=this.state,a=t.selected,n=t.expansions;return p.a.createElement(Te.a,{required:!0,className:e.select},p.a.createElement("select",{multiple:!0,hidden:!0,id:"selectedExpansions",name:"selectedExpansions"},a.map(function(e){return p.a.createElement("option",{selected:!0,key:e,value:e},e)})),p.a.createElement(Ge.a,null,"Expansions"),p.a.createElement(Fe.a,{multiple:!0,displayEmpty:!0,value:a,onChange:this.handleChangeSelect,renderValue:function(t){return p.a.createElement("div",{className:e.chips},t.map(function(t){return p.a.createElement(ze.a,{key:t,label:t,className:e.chip})}))},MenuProps:He},n.map(function(e){return p.a.createElement(A.a,{key:e,value:e},e)})))}},{key:"componentWillMount",value:function(){var e=this;ie.a.get("game/available-expansions").then(function(t){return e.setState(Object(K.a)({},e.state,{expansions:t.data}))})}}]),t}(m.Component),_e=Object(u.withStyles)(function(e){return{select:{width:"100%",minHeight:36},chips:{display:"flex",flexWrap:"wrap"},chip:{margin:e.spacing.unit/4}}})(Ue),qe=a(85),Je=a.n(qe),Ve=a(35),Ye=function(e){return p.a.createElement(Z.a,{variant:"contained",color:"primary",type:"submit",className:e},"Start game")},Xe=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),o=0;o<n;o++)r[o]=arguments[o];return(a=Object(c.a)(this,(e=Object(s.a)(t)).call.apply(e,[this].concat(r)))).state={gameID:a.props.gameID,expansions:[],handSize:10,randomFirstCzar:!0},a.handleSubmit=function(e){return e.preventDefault(),a.props.enoughPlayers?(console.log("Starting game with options",a.state),ie.a.post(T,a.state).catch(function(e){return a.props.pushError(e)})):console.error("Tried to start a game without enough players"),!1},a.handleHandSizeChange=function(e){var t=parseInt(e.target.value);t=Math.min(Math.max(t,0),30),a.setState(Object(K.a)({},a.state,{handSize:t}))},a.handleExpansionSelected=function(e){a.setState(Object(K.a)({},a.state,{expansions:e}))},a}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this.props,t=e.classes,a=e.enoughPlayers,n=this.state,r=n.handSize,i=n.randomFirstCzar;return p.a.createElement("div",{className:t.container},p.a.createElement("form",{className:t.form,onSubmit:this.handleSubmit},p.a.createElement(J.a,{variant:"h6",className:t.formLabel},"Game Options"),p.a.createElement(Te.a,{fullWidth:!0,margin:"normal"},p.a.createElement(_e,{onSelectedChange:this.handleExpansionSelected})),p.a.createElement(Te.a,{required:!0,fullWidth:!0,margin:"normal"},p.a.createElement(Ve.b,{label:"Hand size",id:"handSize",name:"handSize",type:"number",onChange:this.handleHandSizeChange,value:r})),p.a.createElement(Te.a,{fullWidth:!0,margin:"normal"},p.a.createElement(Je.a,{control:p.a.createElement(We.a,{id:"randomFirstCzar",name:"randomFirstCzar",color:"primary",value:i}),label:"First Czar chosen randomly"})),p.a.createElement(le,{className:t.button}),a?p.a.createElement(Ye,{className:t.button}):null))}}]),t}(m.Component),$e=Object(V.b)(function(){},{pushError:oe})(Object(u.withStyles)(function(e){return{container:{padding:2*e.spacing.unit,maxWidth:480,marginLeft:"auto",marginRight:"auto"},formLabel:{textAlign:"left"},form:{marginTop:2*e.spacing.unit,display:"inline-block",width:"100%",textAlign:"right"},button:{margin:e.spacing.unit}}})(Xe)),Ke=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),o=0;o<n;o++)r[o]=arguments[o];return(a=Object(c.a)(this,(e=Object(s.a)(t)).call.apply(e,[this].concat(r)))).updateState=function(){var e=a.props.match.params.gameID;ie.a.get(L(e)).then(function(e){a.setState({room:e.data}),"Not started"===e.data.phase&&window.setTimeout(a.updateState,5e3)}).catch(function(e){return a.props.pushError(e)})},a}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this.props,t=e.classes,a=e.username;if(null==this.state)return p.a.createElement(me.a,{className:t.loading});var n=this.state.room;if("Not started"!==n.phase)return p.a.createElement(Pe,{stateID:n.stateID});var r=n.players.length>2,i=n.owner===a;return p.a.createElement("div",{className:t.container},p.a.createElement(J.a,{variant:"h4",gutterBottom:!0},n.name),r?p.a.createElement(J.a,{variant:"h6",gutterBottom:!0},"Waiting for the game creator to start the game"):p.a.createElement(J.a,{variant:"h6",gutterBottom:!0},"Waiting for more players to join"),p.a.createElement(J.a,null,"Creator: ",n.owner,"."),p.a.createElement(J.a,{gutterBottom:!0},"Players: ",n.players.join(", "),"."),i?p.a.createElement($e,{gameID:n.id,enoughPlayers:r}):p.a.createElement(le,{className:t.button}))}},{key:"componentWillMount",value:function(){this.updateState()}}]),t}(m.Component),Qe=Object(V.b)(function(e){return{username:e.username}},{pushError:oe})(Object(u.withStyles)(function(e){return{container:{textAlign:"center",marginTop:2*e.spacing.unit},button:{margin:e.spacing.unit},loading:{margin:e.spacing.unit}}})(Ke)),Ze=Object(u.withStyles)(function(e){return{linkContainer:{display:"flex",alignItems:"center",justifyContent:"center"},icon:{margin:e.spacing.unit}}})(function(e){var t=e.classes;return p.a.createElement("div",{className:t.linkContainer},p.a.createElement(J.a,null,p.a.createElement("a",{target:"blank",href:"https://github.com/J4RV"},p.a.createElement(O,{className:t.icon}))),p.a.createElement(J.a,null,p.a.createElement("a",{target:"blank",href:"https://store.cardsagainsthumanity.com"},p.a.createElement(H.a,{className:t.icon}))))}),et=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),o=0;o<n;o++)r[o]=arguments[o];return(a=Object(c.a)(this,(e=Object(s.a)(t)).call.apply(e,[this].concat(r)))).state={username:"",password:"",disabled:!1},a.handleChangeUser=function(e){var t=Object.assign({},a.state);t.username=e.target.value.trim(),a.setState(t)},a.handleChangePass=function(e){var t=Object.assign({},a.state);t.password=e.target.value.trim(),a.setState(t)},a.handleSubmit=function(e,t){e.preventDefault(),a.setState(Object(K.a)({},a.state,{disabled:!0}));var n={username:a.state.username,password:a.state.password};return ie.a.post(t,n).then(function(e){return a.props.onSubmitResponse(e)}).catch(function(e){a.setState(Object(K.a)({},a.state,{disabled:!1})),a.props.onError(e.response.data)}),!1},a}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this,t=this.props.classes;return p.a.createElement("div",{className:t.container},p.a.createElement(J.a,{variant:"h2",gutterBottom:!0},"Cards Against Humanity"),p.a.createElement(J.a,{variant:"h4",gutterBottom:!0},"A party game for horrible people."),p.a.createElement("form",{className:t.form,onSubmit:function(t){return e.handleSubmit(t,"user/login")}},p.a.createElement(ge,{isBlack:!0,text:"I'm _ and my password is _.",expansion:"Security questions"}),p.a.createElement(ne.a,{required:!0,fullWidth:!0,margin:"normal",label:"Username",autoComplete:"username",onChange:this.handleChangeUser}),p.a.createElement(ne.a,{required:!0,fullWidth:!0,margin:"normal",label:"Password",type:"password",autoComplete:"password",onChange:this.handleChangePass}),p.a.createElement(Te.a,{margin:"normal",fullWidth:!0},p.a.createElement(Z.a,{type:"submit",variant:"contained",color:"primary",onClick:function(t){return e.handleSubmit(t,"user/login")},disabled:this.state.disabled},"Log in")),p.a.createElement(Te.a,{margin:"normal",fullWidth:!0},p.a.createElement(Z.a,{type:"button",variant:"contained",onClick:function(t){return e.handleSubmit(t,"user/register")},disabled:this.state.disabled},"Register")),p.a.createElement(Ze,null)))}}]),t}(m.Component),tt=function(e){function t(){return Object(i.a)(this,t),Object(c.a)(this,Object(s.a)(t).apply(this,arguments))}return Object(l.a)(t,e),Object(o.a)(t,[{key:"componentWillMount",value:function(){var e=this;ie.a.get("user/valid-cookie").then(function(t){return e.props.processLoginResponse(t)}).catch(function(t){return e.props.processLoginResponse(t)})}},{key:"render",value:function(){var e=this.props,t=e.validCookie,a=e.processLoginResponse,n=e.pushError,r=e.classes;return null==t?p.a.createElement(me.a,null):t?this.props.children:p.a.createElement(et,{onSubmitResponse:a,onError:n,classes:r})}}]),t}(m.Component),at=Object(V.b)(function(e){return{validCookie:e.validCookie}},{processLoginResponse:function(e){return{type:"PROCESS_LOGIN_RESPONSE",payload:{response:e}}},pushError:oe})(Object(u.withStyles)(function(e){return{container:{textAlign:"center",marginTop:2*e.spacing.unit},form:{maxWidth:260,marginTop:2*e.spacing.unit,marginBottom:2*e.spacing.unit,padding:2*e.spacing.unit,display:"inline-block"}}})(tt)),nt=a(87),rt=a.n(nt),it=a(89),ot=a.n(it),ct=a(38),st=a.n(ct),lt=a(88),ut=a.n(lt),mt=a(64),pt=a.n(mt),dt=function(e){var t=e.username,a=e.game,n=e.joinGame;return a.players.includes(t)?p.a.createElement(S.a,{to:"/game/room/".concat(a.id)},p.a.createElement(Ve.a,{color:"primary",variant:"contained"},"Enter")):p.a.createElement(Ve.a,{color:"primary",variant:"contained",onClick:function(){return n(a.id)}},"Join")},ht=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),o=0;o<n;o++)r[o]=arguments[o];return(a=Object(c.a)(this,(e=Object(s.a)(t)).call.apply(e,[this].concat(r)))).state={games:void 0,joinedGame:!1},a.refreshGames=function(){ie.a.get(a.props.fetchGamesUrl).then(function(e){a.setState(Object(K.a)({},a.state,{games:e.data}))}).catch(function(e){return a.props.pushError(e)})},a.joinGame=function(e){ie.a.post(z,{id:e}).then(a.setState(Object(K.a)({},a.state,{joinedGame:e}))).catch(function(e){return a.props.pushError(e)})},a}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this;if(this.state.joinedGame)return p.a.createElement(g.a,{to:"/game/room/".concat(this.state.joinedGame)});var t=this.props.username,a=this.state.games;return null==a?p.a.createElement(me.a,null):0===a.length?p.a.createElement(J.a,{variant:"h6",align:"center"},"No games found"):p.a.createElement(rt.a,null,p.a.createElement(ut.a,null,p.a.createElement(pt.a,null,p.a.createElement(st.a,{align:"left"},"Name"),p.a.createElement(st.a,{align:"center"},"Owner"),p.a.createElement(st.a,{align:"center"},"Actions"))),p.a.createElement(ot.a,null,a.map(function(a){return p.a.createElement(pt.a,{key:a.id},p.a.createElement(st.a,{align:"left"},a.name),p.a.createElement(st.a,{align:"center"},a.owner),p.a.createElement(st.a,{align:"center"},p.a.createElement(dt,{game:a,joinGame:e.joinGame,username:t})))})))}},{key:"componentWillMount",value:function(){this.refreshGames()}}]),t}(m.Component),gt=Object(V.b)(function(e){return{username:e.username}},{pushError:oe})(ht),ft=a(37),bt=a.n(ft),Et=function(e){function t(){return Object(i.a)(this,t),Object(c.a)(this,Object(s.a)(t).apply(this,arguments))}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this.props.classes;return p.a.createElement("div",{className:e.root},p.a.createElement(Ve.c,{variant:"h5",className:e.title},"My games in progress"),p.a.createElement(bt.a,{className:e.tableContainer},p.a.createElement(gt,{fetchGamesUrl:"game/list-in-progress"})))}}]),t}(m.Component),vt=Object(V.b)(null,{pushError:oe})(w()()(Object(u.withStyles)(function(e){return{root:Object(de.a)({maxWidth:960-4*e.spacing.unit,marginTop:3*e.spacing.unit,marginLeft:"auto",marginRight:"auto",overflowX:"auto",padding:2*e.spacing.unit},e.breakpoints.down("sm"),{padding:e.spacing.unit}),tableContainer:{overflowX:"auto"},title:{textAlign:"center",marginBottom:2*e.spacing.unit},createBtn:{float:"right",marginTop:2*e.spacing.unit}}})(Et))),yt=function(e){function t(){return Object(i.a)(this,t),Object(c.a)(this,Object(s.a)(t).apply(this,arguments))}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this.props.classes;return p.a.createElement("div",{className:e.root},p.a.createElement(Ve.c,{variant:"h5",className:e.title},"Open games"),p.a.createElement(bt.a,{className:e.tableContainer},p.a.createElement(gt,{fetchGamesUrl:"game/list-open"})),p.a.createElement(S.a,{to:"/game/list/create"},p.a.createElement(Ve.a,{type:"button",className:e.createBtn},"Create new game")))}}]),t}(m.Component),Ot=Object(V.b)(null,{pushError:oe})(w()()(Object(u.withStyles)(function(e){return{root:Object(de.a)({maxWidth:960-4*e.spacing.unit,marginTop:3*e.spacing.unit,marginLeft:"auto",marginRight:"auto",padding:2*e.spacing.unit},e.breakpoints.down("sm"),{padding:e.spacing.unit}),tableContainer:{overflowX:"auto"},title:{textAlign:"center",marginBottom:2*e.spacing.unit},createBtn:{float:"right",marginTop:2*e.spacing.unit}}})(yt))),jt=a(166),wt=a.n(jt),Ct=function(e){return p.a.createElement(k.a,{key:"close","aria-label":"Close",color:"inherit",onClick:e.onClick,className:e.className},p.a.createElement(wt.a,null))},kt=a(167),St=a.n(kt),xt=a(90),Nt=a.n(xt),Rt=a(63),It=a.n(Rt),Pt=a(169),At=Object(V.b)(function(e){return{errors:e.errors}},{removeError:function(e){return{type:"REMOVE_ERROR",payload:{index:e}}}})(Object(u.withStyles)(function(e){return{error:{color:e.palette.getContrastText(e.palette.error.dark),background:e.palette.error.dark,display:"flex",alignItems:"center"},icon:{marginRight:e.spacing.unit},message:{display:"flex",alignItems:"center"}}})(function(e){var t=e.errors,a=e.children,n=e.removeError,r=e.classes;return p.a.createElement(p.a.Fragment,null,a,p.a.createElement(Nt.a,{anchorOrigin:{vertical:"bottom",horizontal:"left"},open:null!=t&&t.length>0},p.a.createElement(It.a,{className:r.error,message:t.map(function(e,t){return p.a.createElement("div",{className:r.message},p.a.createElement(St.a,{className:r.icon}),e,p.a.createElement(Ct,{onClick:function(){return n(t)}}))})})))})),Wt=a(93),Bt=a(125),zt=a.n(Bt),Lt=a(122),Tt=a.n(Lt),Dt=a(123),Gt=a.n(Dt),Mt={validCookie:void 0,userID:void 0,username:void 0,errors:[]},Ft=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:Mt,t=arguments.length>1?arguments[1]:void 0;switch(t.type){case"PROCESS_LOGIN_RESPONSE":return function(e,t){var a=t.payload.response;if(200!==a.status)return Object(K.a)({},e,{validCookie:!1});var n=a.data;return Object(K.a)({},e,{validCookie:!0,userID:n.id,username:n.username})}(e,t);case"PUSH_ERROR":return function(e,t){var a=t.payload.msg;console.error(a);var n=e.errors.concat(a);return n.length>4&&n.splice(0,1),Object(K.a)({},e,{errors:n})}(e,t);case"REMOVE_ERROR":return function(e,t){var a=t.payload.index,n=Object(Pt.a)(e.errors);return n.splice(a,1),Object(K.a)({},e,{errors:n})}(e,t);default:return e}},Ht=Object(u.createMuiTheme)({palette:{primary:Gt.a,secondary:Tt.a,whitecard:{text:"#161616",background:"#FAFAFA"},blackcard:{text:"#FAFAFA",background:"#161616"},expansion:"#888888",type:"dark"},lights:{glow:"0 0 4px 2px ".concat(zt.a[100],", 0 0 24px 2px ").concat(zt.a[500])}}),Ut=Object(d.a)(function(){return p.a.createElement(At,null,p.a.createElement(at,null,p.a.createElement(Y,null),p.a.createElement(h.a,{exact:!0,path:"/",render:function(){return p.a.createElement(g.a,{to:"/game/list/my-games-in-progress"})}}),p.a.createElement(h.a,{path:"/game/list/create",component:se}),p.a.createElement(h.a,{path:"/game/list/my-games-in-progress",component:vt}),p.a.createElement(h.a,{path:"/game/list/open",component:Ot}),p.a.createElement(h.a,{path:"/game/room/:gameID",component:Qe})))}),_t=function(e){function t(){return Object(i.a)(this,t),Object(c.a)(this,Object(s.a)(t).apply(this,arguments))}return Object(l.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){return p.a.createElement(f.a,null,p.a.createElement(u.MuiThemeProvider,{theme:Ht},p.a.createElement(V.a,{store:Object(Wt.b)(Ft)},p.a.createElement($.a,null),p.a.createElement(Ut,null))))}}]),t}(m.Component),qt=a(30);a.n(qt).a.render(p.a.createElement(_t,null),document.getElementById("root")),function(e){if("serviceWorker"in navigator){if(new URL("",window.location.href).origin!==window.location.origin)return;window.addEventListener("load",function(){var t="".concat("","/service-worker.js");n?(function(e,t){window.fetch(e).then(function(a){var n=a.headers.get("content-type");404===a.status||null!=n&&-1===n.indexOf("javascript")?navigator.serviceWorker.ready.then(function(e){e.unregister().then(function(){window.location.reload()})}):r(e,t)}).catch(function(){console.log("No internet connection found. App is running in offline mode.")})}(t,e),navigator.serviceWorker.ready.then(function(){console.log("This web app is being served cache-first by a service worker. To learn more, visit http://bit.ly/CRA-PWA")})):r(t,e)})}}()}},[[230,2,1]]]);
//# sourceMappingURL=main.17c9a2eb.chunk.js.map