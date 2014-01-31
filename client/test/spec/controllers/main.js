'use strict';

describe('Controller: MainCtrl', function () {

  // load the controller's module
  beforeEach(module('clientApp'));

  var MainCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    MainCtrl = $controller('MainCtrl', {
      $scope: scope
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    //expect(scope.awesomeThings.length).toBe(3);
  });
});

describe('Controller: NavCtrl', function () {
  beforeEach(module('clientApp'));

  var ctrl, scope, rootScope, location;

  beforeEach(inject(function ($location, $controller, $rootScope) {
    location = $location;
    rootScope = $rootScope;
    scope = $rootScope.$new();
    ctrl = $controller('NavCtrl', {
      $scope: scope
    });
  }));

  it('should say when we are on a page', function () {
    location.path('/');
    rootScope.$apply();
    expect(scope.isActive('/')).toBe(true);

    location.path('/contact');
    rootScope.$apply();
    expect(scope.isActive('/about')).toBe(false);
  });
});
