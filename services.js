app.service('vendingservice', function ($http) {

    
    this.post = function (api, student) {
        var request = $http({
            method: "post",
            url: api,
            data: student,
          //  headers:{'content-Type':'application/json'},
        });
        return request;
    }


    this.getall = function (api) {

        url = api;
        return $http.get(url);
    }
    this.put=function(api,dta,id){
        var request=$http({method:"put",url:api+id,data:dta});
        return request;
    }
    this.getbyid=function(api){
        url=api;
return $http.get(url);
    }
});