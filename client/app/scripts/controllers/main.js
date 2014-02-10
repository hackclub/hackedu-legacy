'use strict';

angular.module('hackedu')
  .controller('MainCtrl', function ($scope) {
    $scope.map = {
      center: {
        latitude: 36,
        longitude: -101
      },
      zoom: 4
    };
    
    $scope.schools = [
      {
        'name': 'Austin High School',
        'location': {
          'latitude': '30.27382',
          'longitude': '-97.76745'
        }
      },
      {
        'name': 'Thunderridge High School',
        'location': {
          'latitude': '39.5347968',
          'longitude': '-105.01200670000003'
        }
      }
    ];
  })
  .controller('NavCtrl', function ($scope, $location) {
    $scope.isActive = function (viewLocation) {
      return viewLocation === $location.path();
    };
  })
  .controller('ApplyCtrl', function ($scope, $location) {
    $scope.isActive = function (viewLocation) {
      return viewLocation === $location.path();
    };
  });

angular.module('hackedu')
  .directive('match', function($parse) {
    return {
      require: 'ngModel',
      link: function(scope, elem, attrs, ctrl) {
        scope.$watch(function() {
          return $parse(attrs.match)(scope) === ctrl.$modelValue;
        }, function(currentValue) {
          ctrl.$setValidity('mismatch', currentValue);
        });
      }
    };
  });
