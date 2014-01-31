'use strict';

angular.module('clientApp')
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
  });

angular.module('clientApp')
  .controller('NavCtrl', function ($scope, $location) {
    $scope.isActive = function (viewLocation) {
      return viewLocation === $location.path();
    };
  });
