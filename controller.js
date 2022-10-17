
  var  app = angular.module('myapp', []);



  app.controller('vendingcontroller', ['$scope','vendingservice', function($scope,vendingservice) {
    var baseurl = 'http://localhost:8080/';
    $scope.inicash = 0;
    $scope.slotcsh = [20,50,100,200,150];

  // it runs when user clicked the cancel
  $scope.cancel = function(){$scope.inicash="";}
  
    //for incrementing cash when user add more cash
    function increcash(v) {     
      $scope.inicash += v;
      console.log('cash deposited by user',$scope.inicash)
    }
//for decrementing cash when user purchased product
function decrecash(v) {
  $scope.inicash -= v;
}



    // when cash inserted and get incremented by adding more cash 
    $scope.insertcoin = function(v) {
      increcash(v);
    }
    // when user selected the product
    $scope.p = "20";
    $scope.pro="maza";
    $scope.selectslot = function(slot) {
      if($scope.inicash < price){
alert("Insufficient cash")
return false;
      }
       if(ma[0]=="0"){
        alert("out of stock")
        return false;
      }   
      var price = getsltprice(slot);//price of product by slot
     console.log('price of prodt in slot',price)
      if($scope.inicash>= price){   //initial cash >= to price of product
        var dta = {
          name: $scope.p,
          price: $scope.pro, 
      }
        var api = baseurl + 'saveproduct/';
        var pst = vendingservice.post(api, dta);
        pst.then(function (response){
          if (response.data != ""){
          }
          decrecash(price);     
        })       
$scope.up = function(slot) {
  var price = getsltprice(slot);
var dta = {
  id:$scope.id, 
}
var price = getsltprice(slot);
console.log('price of prodt in slot',price)
var api = baseurl + 'uptproduct/';
var pst = vendingservice.put(api, dta,$scope.id);
pst.then(function (response){ 
    alert("purchasd successfully")
    $scope.getbyid();
    }        
)

//decrecash(price);
//$scope.getbyid();
//setTimeout(function(){
  //location.reload(); 
//}, 2000); 
}      
    }  
     else{
      alert("insufficient cash")  
    }   
  }

 // $scope.getbyid=function(){$route.reload();}


    $scope.pr = "50";
    $scope.prod="pepsi";
    $scope.selects = function(slot) {
      var price = getsltprice(slot);
     if($scope.inicash>= price){ 
        var dta = {         
          name: $scope.pr,
          price: $scope.prod, 
      }
        var api = baseurl + 'saveproduct2/';
        var pst = vendingservice.post(api, dta);
        pst.then(function (response){
          if (response.data != ""){
            alert("purchased successfully")
          }       
        })         
      decrecash(price);     
    }  else{
      alert("insufficient cash")
    }}

    $scope.pri = "100";
    $scope.produ="cocacola";
    $scope.selectsl = function(slot) {
      var price = getsltprice(slot);
     if ($scope.inicash >= price){
        var dta = {
          name: $scope.pri,
          price: $scope.produ, 
      }
        var api = baseurl + 'saveproduct3/';
        var pst = vendingservice.post(api, dta);
        pst.then(function (response){
          if (response.data != ""){
            alert("purchased successfully")
          }          
        })        
      decrecash(price);     
    }else{
      alert("insufficient cash")
    }}
   
    $scope.price = "200";
    $scope.product="sprite";
    $scope.selectslo = function(slot) {
      var price = getsltprice(slot);
     if ($scope.inicash >= price){    
        var dta = {
          name: $scope.price,
          price: $scope.product, 
      }
        var api = baseurl + 'saveproduct4/';
        var pst = vendingservice.post(api, dta);
        pst.then(function (response){
          if (response.data != ""){
            alert("purchased successfully")
          }
          
        })         
      decrecash(price);     
   }else{
    alert("insufficient cash")
  }}

  $scope.pricee = "150";
  $scope.products="redbull";
  $scope.selectslots = function(slot) {
    var price = getsltprice(slot);
    console.log("red",price)
   if ($scope.inicash >= price){    
      var dta = {
        name: $scope.pricee,
        price: $scope.products, 
    }
      var api = baseurl + 'saveproduct5/';
      var pst = vendingservice.post(api, dta);
      pst.then(function (response){
        if (response.data != ""){
          alert("purchased successfully")
        }       
      })          
    decrecash(price);   
    qty();
 }else{
  alert("insufficient cash")
}}

function qty() {
  var api = baseurl + 'getquantity/';
  //$scope.qty=[]; 
  var get = vendingservice.getall(api);
  get.then(function (response) {   
    $scope.qty=response.data;     
    console.log('quantities',$scope.qty)         
  });
}
qty();
$scope.getbyid=function(){
  var api=baseurl+'getbyid/';
  var get=vendingservice.getbyid(api);
  get.then(function (response){
    $scope.qtym=response.data;
    //console.log("User = " + JSON.stringify(response));
   // console.log(response)
    var c=JSON.stringify(response)
    var d=c[50]
    ma[0]=d;
    //console.log(d)
    return d;
  })
}
var ma = [];
console.log(ma)
$scope.getbyid();
//var a=document.getElementById("qty");

var s= $scope.qty;
//JSON.stringify(s)
console.log("qty value",s)
// getting price of slot
    getsltprice= function (slot) {
      console.log('slot no',slot)
      return $scope.slotcsh[slot];
    }
  
    // for showing products,when user inserted the cash
    //function viewpr() {
     //   var dvSites = document.querySelector('#dvSites');          
      //  dvSites.style.display = '';
   // }
    //document.querySelector('#btnshow').addEventListener('click', viewpr);  
  }]);
