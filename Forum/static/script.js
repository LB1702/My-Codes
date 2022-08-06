console.log("charged")

function init(){
  let cookies = document.cookie;
  console.log(cookies)
  if (cookies.length != 0){
    let buttonSignIn = document.getElementById("register");
    let buttonLogIn = document.getElementById("login");
    // console.log(buttonLogIn)
    buttonLogIn.remove()
    buttonSignIn.remove()
    // console.log(window.location.pathname)
    if(window.location.pathname != '/profile'){
      let nameValue = document.getElementById("profile").attributes.getNamedItem('value').value;
      let buttonProfile = document.getElementById("profile").querySelector('a')
      buttonProfile.innerHTML = nameValue
    }

    // console.log(buttonProfile.querySelector('a'))
    let menu = document.getElementById("menu")
    let butttonDeconnect = document.createElement('button')
    butttonDeconnect.innerHTML = 'Log Out'
    butttonDeconnect.setAttribute('onclick', 'deconnexion()')
    menu.appendChild(butttonDeconnect)
    console.log("Done!")
  }
}

function deconnexion(){
  console.log("decooooo")
  document.cookie.split(";").forEach(function(c) { document.cookie = c.replace(/^ +/, "").replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/"); });
  location.reload()
}

const criteria = ["• at least 8 characters <br>",
                   "• at least 1 uppercase letter <br> ",
                   "• at least 1 digit <br> ",
                   "• at least 1 special character"];
let securityLevel0 = false;
let securityLevel1 = false;
let securityLevel2 = false;
let securityLevel3 = false;

function passCriteria(){
    let span = document.querySelector('#tooltiptext');
    span.style.width = '250px';
    span.style.height = '120px';
    let introCriteria = "Password must have: <br> <br>" ; 
    let filler = '';
    let count = 0;
    criteria.forEach(element => {
        filler += `<span id="criteria-${count.toString()}" style='color: red;'>${element}</span>`;
        count++
    });
    span.innerHTML = introCriteria + filler;

}

function changeCriteriaColor(id, color){
    // console.log("Changing color...")
    let strId = id.toString()
    let span = document.getElementById('criteria-' + strId);
    span.style.color = color;

}

function checkPassword(){
    passCriteria()
    let exceptions = ['Shift', 'CapsLock','AltGraph', 'Control',
    'ArrowLeft','ArrowRight','ArrowDown','ArrowUp']
    let passField = document.getElementById('passwd');
    let password = ""
    passField.addEventListener("keyup", event => {
        if (event.isComposing || event.keyCode === 229) {
          return;
        }
        let inputText = event.key
        if (inputText == 'Backspace'){
            password = password.substring(0, password.length - 1);
        } else if (exceptions.includes(inputText)){ 
            console.log('pressed a function key')
        }else{
            password += inputText
        }
        
        let strength = 0
        // See what's in password field
        // console.log(password)
        if (password.match(/[A-Z]+/)) {
            changeCriteriaColor(1, 'green')
            securityLevel1 = true

          } else{
              changeCriteriaColor(1, 'red')
            securityLevel1 = false

          }
          if (password.match(/[0-9]+/)) {
            changeCriteriaColor(2, 'green')
            securityLevel2 = true

          } else{
              changeCriteriaColor(2, 'red')
            securityLevel2 = false

          }
          if (password.match(/[$@#&!]+/)) {
            changeCriteriaColor(3, 'green')
            securityLevel3 = true
            
          } else{
              changeCriteriaColor(3, 'red')
            securityLevel3 = false

          }

          if (password.length > 8) {
            changeCriteriaColor(0, 'green')
            securityLevel0 = true
          } else {
            changeCriteriaColor(0, 'red')
            securityLevel0 = false

          }
// console.log(securityLevel0, securityLevel1, securityLevel2, securityLevel3)

      });

}




function clearSpan(){
    document.querySelector('#tooltiptext').innerHTML = "";
}

function checkReg() {
  let criteriasToCheck = [securityLevel0, securityLevel1, securityLevel2, securityLevel3];
  criteriasToCheck.forEach(element => {
    if (element === false){
      alert('Attention! One or more pass criteria are not satisfied.')
      document.getElementById('formReg').action = '/form.html'
      Location.reload()
    } 
  });
  
}

function setCookie(name,value,days) {
  var expires = "";
  if (days) {
      var date = new Date();
      date.setTime(date.getTime() + (days*24*60*60*1000));
      expires = "; expires=" + date.toUTCString();
  }
  document.cookie = name + "=" + (value || "")  + expires + "; path=/";
}


function redirectToHomePage(){
  location.replace("/index")
}


function Sorry(){
  const message = "Desole faut accepter frerot"

  resultat = window.confirm(message);

  if (resultat) {
    deleteCookie('loop')
    location.replace("/accueil");
} else {
  setCookie('loop', 'true', 1)
  location.reload()
}
}

function deleteCookie(name) {
  document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}



function SorryAccueil(){
  let cookies = document.cookie
  let loopCheck = cookies.split('=')
  if (loopCheck[0] === 'loop' && loopCheck[1] === 'true'){
    location.replace('/accueil')
  }
}


function SpanPostCreation(){
  document.getElementById("postCreation").style.visibility = 'visible';
  return
}

function ClearSpanPostCreation(){
  document.getElementById("postCreation").style.visibility = 'hidden';
  return
}

const toBase64 = file => new Promise((resolve, reject) => {
  const reader = new FileReader();
  reader.readAsDataURL(file);
  reader.onload = () => resolve(reader.result);
  reader.onerror = error => reject(error);
});


function getDataUrl(img) {
  // Create canvas
  const canvas = document.createElement('canvas');
  const ctx = canvas.getContext('2d');
  // Set width and height
  canvas.width = img.width;
  canvas.height = img.height;
  // Draw the image
  ctx.drawImage(img, 0, 0);
  return canvas.toDataURL('image/jpeg');
}

function convertImageOnLoad(){
  // const img = document.querySelector('#postImage');
  // console.log(img)
  // img.addEventListener('mouseover', function (event) {
  //    const dataUrl = getDataUrl(event.currentTarget);
  //    console.log(dataUrl);
  // });
  // const img = document.querySelector('#blah').files[0];
  // console.log(getDataUrl(img))
 
  var file = document.querySelector(
    'input[type=file]')['files'][0];

var reader = new FileReader();
console.log("next");
  
reader.onload = function () {
    // base64String = reader.result.replace("data:", "")
    //     .replace(/^.+,/, "");

    // imageBase64Stringsep = base64String;
    let imageBase64Formatted = reader.result;
    // alert(imageBase64Stringsep);
    document.getElementById("imageConverted").value = imageBase64Formatted;
    console.log(reader.result);
}
reader.readAsDataURL(file);



}




function showImage(input) {
  if (input.files && input.files[0]) {
    var reader = new FileReader();
    reader.onload = function (e) {
      let imgToLoad = document.querySelector('#blah')
        imgToLoad.src = e.target.result;
        // .width = 150
        // .height = 200;
        imgToLoad.setAttribute('width', 150)
        imgToLoad.setAttribute('height', 150);
        // .width(150)
        // .height(200);
    };
    reader.readAsDataURL(input.files[0]);
  }

  convertImageOnLoad()
}


function thankYouPost(){
  ClearSpanPostCreation()
  spanThankYouMessage()
}


function spanThankYouMessage(){
  document.getElementById("success").style.visibility = 'visible'
}



function assignIdToPostsAndRefs(){
  let posts = document.querySelectorAll("#post")
  // console.log(posts)
  let count = 1;
  posts.forEach( element => {
    element.id = "post-"+ count.toString()
    count++
    // console.log(element)
  })

  let refs = document.querySelectorAll(".seeMoreLink")
  count = 1
  refs.forEach(element => {
    element.id = "seeMoreLink-" + count.toString()
    element.href = '/post/' + count.toString()
    count++
  })
  // console.log(refs)
}

function assignIdToComments(){
  console.log('running...')
  let comments = document.querySelectorAll(".comment")
  // console.log(posts)
  let count = 1;
  comments.forEach( element => {
    element.id = "comment-"+ count.toString()
    count++
    // console.log(element)
  })
}




function tabulatePosts(){
  let posts = document.querySelectorAll(".post")
  // console.log(posts.length)
  let x= 0;  // 0% | 50% OK
  let y = 300;
  done = 0;
  posts.forEach( element => {
    // console.log(element)
    element.style.marginLeft = x.toString() + '%' ;
    element.style.marginTop = y.toString() + 'px' ;
   
    if(done % 2 === 0){
      x = 50
    } else{
      x = 0
      y +=205
    }

    done++
  })

}



function tabulateComments(){
  let comments = document.querySelectorAll(".comment")
  let insertComment = document.querySelector(".insertComment") 

  // console.log(posts.length)
  // let x= 0;  // 0% | 50% OK
  let y = 500;
  done = 0;
  comments.forEach( element => {
    // console.log(element)
    // element.style.marginLeft = x.toString() + '%' ;
    element.style.marginTop = y.toString() + 'px' ;
    y +=180
    done++
  })
  y+= 100
  insertComment.style.marginTop = y.toString() + 'px' ;

}




function liked(){
  var element = document.getElementById("like");
  element.classList.toggle("liked");
}

function regSuccess(){
  window.setTimeout(function(){

    // Move to a new location or you can do something else
    window.location.href = "/login";

}, 3000);

}

function commentSuccess(){
  alert(document.querySelector(".contentID").name)
  window.location.href = "/post";

}