'use strict';

describe('Controller: MainCtrl', function () {
  beforeEach(module('clientApp'));

  var ctrl, scope;

  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ctrl = $controller('MainCtrl', {
      $scope: scope
    });
  }));

  it('should attach a map to the scope', function () {
    expect(scope.map).toNotBe(null);
  });

  it('should attach schools to the scope', function () {
    expect(scope.schools.length).toBe(2);
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
