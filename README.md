# Biznesa vadības sistēma Genos

## Projekta apraksts

Šis ir PIKC "Rīgas Valsts tehnikums" kvalifikācijas darbs biznesa vadības sistēma Genos. Mans mērķis ir izveidot 
iekšējo sistēmu priekš kompānijas, kurā es strādāju, tagad plānotās iespējas ir: Kalendārs(kur var redzēt jau gatavus events un izveidot pats savus),
 kas tiek ērti attēlots lietotājiem, droša login sistēma, lietotāju pārvaldīšanas sistēma, aptauju funkcionalitāte un  laika pārvaldības sistēma 

## Izmantotās tehnoloģijas
Projektā tiek izmantots:
  - Golang
  - Docker
  - Docker-compose
  - Composer
  - MySQL
  - Vuejs
    - Vuex
    - Vue Router 
    - BootstrapVue
  - SCSS
  - Biblotēkas
    - axios (https://github.com/axios/axios)
    - 

## Izmantotie avoti
[W3School](https://www.w3schools.com/html/default.asp) - tika ņemts example html kods.

[The 2019 Frontend Developer Crash Course - HTML & CSS Tutorial for Beginners](https://www.youtube.com/watch?v=8gNrZ4lAnAw) - kā pamats HTML, CSS apguvei skatījos videoklipu un apguvu kā pievienot CSS pie HTML.
  
  - Golang
  
    [codementor](https://www.codementor.io/@tamizhvendan/managing-data-in-golang-using-gorm-part-1-a9cdjb8nb) - kā izveidot login controller
    
    [Medium](https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72) - kā hash paroles golang(izmantoju doto kodu)
    
    [Youtube](https://www.youtube.com/watch?v=Iq2qT0fRhAA) kā veidot api izmantojot fiber un gorm biblotēkas 
    
    [GoByExample](https://gobyexample.com/command-line-arguments) - Cli (izmantoju doto kodu)
    
    [github](https://github.com/gofiber/jwt) - jwt web tokens izmantot priekš auth(izmantoju doto kodu)
    
    [Gorm](https://gorm.io/docs/index.html) - izmantoju docs priekš Gorm ORM lietošanas (izmantots exmaple kods)
    
    [Fiber](https://docs.gofiber.io/) - izmantoju docs priekš servera taisīšanas(izmantots exmaple kods)
    
    [Github](https://github.com/gofiber/recipes/tree/master/auth-jwt) - kā izveidot auth api ar golang izmantojot jwt(izmantoju doto kodu)
  - Docker
  
     [Dev.to](https://dev.to/aschmelyun/the-beauty-of-docker-for-local-laravel-development-13c0) - Kā izveidot docker-compose un vajadzīgās datnes priekš esošā projekts
     
     [Github](https://github.com/aschmelyun/docker-compose-laravel) -  Kā izveidot docker-compose un vajadzīgās datnes priekš esošā projekts
     
     [Digitalocean](https://www.digitalocean.com/community/tutorials/how-to-set-up-laravel-nginx-and-mysql-with-docker-compose) - Kā izveidot docker-compose un vajadzīgās datnes priekš esošā projekts
  
  - Laravel
  
    [Laravel Docs](https://laravel.com/docs/8.x) - izmantots example kods
    
    [Medium](https://medium.com/modulr/create-api-authentication-with-passport-of-laravel-5-6-1dc2d400a7f) - kā uzstādīt backend auth sistēmu izmantojot oauth2
  
    [Medium](https://medium.com/modulr/create-api-authentication-passport-in-laravel-5-6-confirm-account-notifications-part-2-5e221b021f07) - jaunu lietotāju reģistrēšanas ziņa uz epastu  
  
    [Youtube](https://www.youtube.com/watch?v=ojGbaJuQXe8&t=294s) - Uzlikt gmail kā sūtītāju
    
    [Medium](https://medium.com/modulr/api-rest-with-laravel-5-6-passport-authentication-generate-avatar-part-3-d92ec7935eff) - Kā izveidot avatar izmantojot laravolt package
  
    [Youtube](https://www.youtube.com/watch?v=V5xINbA-z9o) - Daudz pret daudz attiecības laravel

    [StackOverflow](https://stackoverflow.com/questions/30507849/counting-a-many-to-many-relationship) - kā iegūt many to many attiecību skaitu
  - Vuejs
    
    [Vuejs Docs](https://vuejs.org/v2/guide/installation.html) - Izmantots example kods
    
    [ValidateJs docs](https://validatejs.org/) - Izmantots example kods
    
    [stackoverflow](https://stackoverflow.com/questions/48402747/nuxt-js-vue-js-setting-axios-auth-token-in-vuex-store-resets-after-refresh) - axios request intercept

    [BootstrapVue](https://bootstrap-vue.org/) - izmantots exmaple kods 
  - Vuejs un Laravel
  
    [Pagination](https://github.com/gilbitron/laravel-vue-pagination) - Izmantots example kods, lai izveidotu pagnāciju 
    
    [youtube](https://www.youtube.com/watch?v=zMk52OlK2Aw&t=146s) - lapu pagnācija
    
    [youtube](https://www.youtube.com/watch?v=C1r85Q3BFqQ&list=PLJpBh2VJhy5wPhAmjDB42pkHUnqolqxxq&index=5 ) - vuejs spa autorazācija ar tokens  izmantojot laravel backend
    
    [stackoverflow](https://stackoverflow.com/questions/49414697/how-to-change-vue-js-data-value-when-screen-size-changes) - mainīt mainīgo vuejs ,kad mainās ekrāna izmērs    
 
  - CSS
  
    [Material colors](https://material.io/resources/color/#!/?view.left=0&view.right=0&primary.color=FFCCBC) - krāsu izvēle
    
    [Fonts](https://fonts.google.com/?sidebar.open=true&selection.family=Roboto) - Fonti
    
    [Ilustrations](https://www.pixeltrue.com/free-illustrations)
    
    [bootstrap](https://getbootstrap.com/) - Css stils pāris lietām
    
    [W3school](https://www.w3schools.com/css/default.asp) - css stila lietas
    
  - HTML
  
    [Icons](https://tablericons.com/) - svg ikonas
  
:information_source: :exclamation: *Obligāti jānorāda visi avoti, kas ir izmantoti, kas ir skatīts, kāds kods ir kopēts, ja tas netiks norādīts un tiks identificēts, ka kods ir kopēts darbs netiks ieskaitīts.* :exclamation:

## Uzstādīšanas instrukcijas
