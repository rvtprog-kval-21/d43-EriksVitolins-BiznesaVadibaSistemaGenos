# Biznesa vadības sistēma Genos

## Projekta apraksts

Šis ir PIKC "Rīgas Valsts tehnikums" kvalifikācijas darbs biznesa vadības sistēma Genos. Mans mērķis ir izveidot 
iekšējo sistēmu priekš kompānijas, kurā es strādāju, tagad plānotās iespējas ir: Kalendārs(kur var redzēt jau gatavus events un izveidot pats savus),
 kas tiek ērti attēlots lietotājiem, droša login sistēma, lietotāju pārvaldīšanas sistēma, aptauju funkcionalitāte un  laika pārvaldības sistēma, blogs, kalendārs, minimāla projektu vadītbas sistēma, kompānijas forums, 

## Izmantotās tehnoloģijas
Projektā tiek izmantots:
  - Golang
    - Gin
    - Gorm
  - Docker
  - Docker-compose
  - MySQL
  - Vuejs
    - Vuex
    - Vue Router 
    - BootstrapVue
    - vue-tags-input
    - validate.js
    - vue-good-table
    - vue-select
    - vue-showdown
    - vue2-datepicker
    - vue2-timepicker
    - vue-simple-calendar
    - vue-month-picker
    - vue-form-wizard
    - vue-quill-editor
  - SCSS

## Izmantotie avoti
  - Golang
  
    [codementor](https://www.codementor.io/@tamizhvendan/managing-data-in-golang-using-gorm-part-1-a9cdjb8nb) - kā izveidot login controller
    
    [Medium](https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72) - kā hash paroles golang(izmantoju doto kodu)
    
    [Youtube](https://www.youtube.com/watch?v=Iq2qT0fRhAA) kā veidot api izmantojot fiber un gorm biblotēkas 
    
    [GoByExample](https://gobyexample.com/command-line-arguments) - Cli (izmantoju doto kodu)
    
    [github](https://github.com/gofiber/jwt) - jwt web tokens izmantot priekš auth(izmantoju doto kodu)
    
    [Gorm](https://gorm.io/docs/index.html) - izmantoju docs priekš Gorm ORM lietošanas (izmantots exmaple kods)
    
    [Gin](https://github.com/gin-gonic/gin) - izmantoju docs priekš servera taisīšanas(izmantots exmaple kods)
    
    [Github](https://github.com/gofiber/recipes/tree/master/auth-jwt) - kā izveidot auth api ar golang izmantojot jwt(izmantoju doto kodu)
  
    [Stackoverflow](https://stackoverflow.com/questions/54665442/cors-doesnt-work-on-gin-and-golang-group-routes) - kods saistībā ar cors
    
    [Github](https://github.com/gin-gonic/contrib) middleware izmantošana ar tonic, piemēram jwt auth 
 
    [Gin](https://gin-gonic.com/docs/examples) Gin Ietvera example kods 
    
    [Medium](https://medium.com/wesionary-team/sending-emails-with-go-golang-using-smtp-gmail-and-oauth2-185ee12ab306) - epastu sūtīšana uz golang
 
    [youbasics](https://yourbasic.org/golang/maps-explained/) - maps iekš golang
    
    [Article](https://www.loginradius.com/blog/async/sending-emails-with-golang/) - ēpastu sūtīšana
    
    [Go](https://golang.org/doc/) - docs priekš golang
  
    [stackoverflow](https://stackoverflow.com/questions/59249418/how-to-get-the-claims-from-the-token) - kā parse jwt
    
    [Github](https://github.com/AnupKumarPanwar/Golang-realtime-chat-rooms) - kā izveidot real time chat
    
    [Github](https://github.com/gorilla/websocket/issues/367) - kā salabot orgin not allowed

    []

    [stackoverflow](https://stackoverflow.com/questions/37932551/mkdir-if-not-exists-using-golang) - kā izveidot mapes, ja tās neeksistē
  - Docker
  
     [Dev.to](https://dev.to/aschmelyun/the-beauty-of-docker-for-local-laravel-development-13c0) - Kā izveidot docker-compose un vajadzīgās datnes priekš esošā projekts
     
     [Github](https://github.com/aschmelyun/docker-compose-laravel) -  Kā izveidot docker-compose un vajadzīgās datnes priekš esošā projekts
     
     [Digitalocean](https://www.digitalocean.com/community/tutorials/how-to-set-up-laravel-nginx-and-mysql-with-docker-compose) - Kā izveidot docker-compose un vajadzīgās datnes priekš esošā projekts
  
  - Vuejs
    
    [Vuejs Docs](https://vuejs.org/v2/guide/installation.html) - Izmantots example kods
    
    [ValidateJs docs](https://validatejs.org/) - Izmantots example kods

    [Github](https://github.com/phoenixwong/vue2-timepicker) - Izmantots example kods
    
    [Docs](https://binarcode.github.io/vue-form-wizard/#/) - Izmantots example kods
    
    [NPM](https://www.npmjs.com/package/vue-month-picker) - Izmantots example kods
    
    [NPM](https://www.npmjs.com/package/vue-simple-calendar) - Izmantots example kods
    
    [Stackoverflow](https://stackoverflow.com/questions/34722862/how-do-i-remove-time-part-from-javascript-date) - Kā noņemt datumu
    
    [NPM](https://www.npmjs.com/package/vue2-datepicker) - Izmantots example kods
    
    [stackoverflow](https://stackoverflow.com/questions/48402747/nuxt-js-vue-js-setting-axios-auth-token-in-vuex-store-resets-after-refresh) - axios request intercept

    [BootstrapVue](https://bootstrap-vue.org/) - izmantots exmaple kods 

    [Bubble Sort](https://stackabuse.com/merge-sort-in-javascript/) - kā izmantot bubble sort

    [Stackoverflow](https://stackoverflow.com/questions/492994/compare-two-dates-with-javascript) - kā salīdzināt laiku

    [SomeTutorial](https://tutorialedge.net/javascript/vuejs/vuejs-websocket-tutorial/) - kā izveidot websocket connection
    
    [stackoverflow](https://stackoverflow.com/questions/49414697/how-to-change-vue-js-data-value-when-screen-size-changes) - mainīt mainīgo vuejs ,kad mainās ekrāna izmērs  
  
    [NPM](https://www.npmjs.com/package/vue-chat-scroll) - kā izmantot libary

    [youtube](https://www.youtube.com/watch?v=C1r85Q3BFqQ&list=PLJpBh2VJhy5wPhAmjDB42pkHUnqolqxxq&index=5 ) - vuejs spa autorazācija ar tokens
      
    [Stackoverflow](https://stackoverflow.com/questions/1643320/get-month-name-from-date/18648314#18648314) - datumu formatēšana
    
    [Vue select](https://vue-select.org/guide/options.html#options-prop) - vue select docs, izmantoju piemērus
    
    [Vue table](https://xaksis.github.io/vue-good-table/guide/configuration/#columns) - vue table docs, izmantoju piemērus
    
    [Vue markdown](https://vue-showdown.js.org/guide/#npm) - vue markdown docs, izmantoju piemērus
    
    [Stackoverflow](https://stackoverflow.com/questions/13146418/find-all-the-days-in-a-month-with-date-object) - dabūt visas dienas mēnesī   

    [Stackoverflow](https://stackoverflow.com/questions/1643320/get-month-name-from-date/18648314#18648314) - formtatēt datumus

    [Vue Rich text editor](https://github.com/surmon-china/vue-quill-editor) - vue Rich text editor docs, izmantoju piemērus
  
    [Github](https://github.com/richardtallent/vue-simple-calendar) - tika izmantots example kods priekš kalendārs
    
    [Github](https://github.com/richardtallent/vue-simple-calendar-sample) - tika izmantots example kods priekš kalendārs
  - CSS
  
    [Material colors](https://material.io/resources/color/#!/?view.left=0&view.right=0&primary.color=FFCCBC) - krāsu izvēle
    
    [Fonts](https://fonts.google.com/?sidebar.open=true&selection.family=Roboto) - Fonti
    
    [Ilustrations](https://www.pixeltrue.com/free-illustrations)
    
    [bootstrap](https://getbootstrap.com/) - Css stils pāris lietām
    
    [W3school](https://www.w3schools.com/css/default.asp) - css stila lietas
    
    [link](https://demos.creative-tim.com/bootstrap-vue-argon-dashboard-pro/#/!) - dizains
    
    [codepen](https://codepen.io/rikschennink/pen/mjywQb) - burbuļa dizains un kopēts kods
    
  - HTML
  
    [Icons](https://tablericons.com/) - svg ikonas
  

## Uzstādīšanas instrukcijas
Vajadzīgais software: docker-compose, Go, npm 

1. git clone https://github.com/rvtprog-kval-21/d43-EriksVitolins-BiznesaVadibaSistemaGenos

2. cd d43-EriksVitolins-BiznesaVadibaSistemaGenos

3. dokcer-compose up -d --build

4. cd golang_api

5. go run main.go migrate

6. go run main.go server

7. cd ..

8. cd web

9. npm run serve
