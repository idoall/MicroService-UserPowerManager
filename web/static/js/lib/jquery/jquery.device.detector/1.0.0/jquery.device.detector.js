/**
 * @fileOverview device detector
  jQuery Plugin to get Device and Browser informations
 * @author Simon Gattner <npm@0x38.de>
 * @license MIT
 * @version 1.0.0
 */

/**
 * @external "jQuery.fn"
 * @see {@link http://docs.jquery.com/Plugins/Authoring The jQuery Plugin Guide}
 */

(function($) {
  'use strict';
  /**
   * jQuery Methods to get Device and Browser informations
   * @function external:"jQuery.fn".deviceDetector
   * @external "jQuery.fn.deviceDetector"
   */

  $.fn.deviceDetector = function(options) {
    if (typeof options !== undefined) $.extend(true, config, options);
    return true;
  };

  /**
   * Method to detect mobile devices.
   * @function external:"jQuery.fn.deviceDetector".isMobile
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isMobile = function() {
    return isDeviceType(
      config.vendors.apple.ios.pattern.include,
      config.vendors.apple.ios.pattern.exclude
    ) ||
    isDeviceType(
      config.vendors.google.android.pattern.include,
      config.vendors.google.android.pattern.exclude
    ) ||
    isDeviceType(
      config.vendors.microsoft.windowsPhone.pattern.include,
      config.vendors.microsoft.windowsPhone.exclude
    ) ||
    isDeviceType(
      config.vendors.blackberry.blackberry.pattern.include,
      config.vendors.blackberry.blackberry.exclude
    ) ||
    isDeviceType(
      config.mobile.pattern.include,
      config.mobile.pattern.exclude
    );
  };

  /**
   * Method to detect desktop devices.
   * @function external:"jQuery.fn.deviceDetector".isDesktop
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isDesktop = function() {
    return isDeviceType(
      config.vendors.apple.macos.pattern.include,
      config.vendors.apple.macos.pattern.exclude
    ) ||
    isDeviceType(
      config.vendors.lbu.unixlike.pattern.include,
      config.vendors.lbu.unixlike.pattern.exclude
    ) ||
    isDeviceType(
      config.vendors.microsoft.windows.pattern.include,
      config.vendors.microsoft.windows.pattern.exclude
    );
  };

  /**
   * Method to detect Safari.
   * @function external:"jQuery.fn.deviceDetector".isSafari
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isSafari = function() {
    return isDeviceType(
      config.vendors.apple.safari.pattern.include,
      config.vendors.apple.safari.pattern.exclude
    );
  };

  /**
   * Method to detect iPad.
   * @function external:"jQuery.fn.deviceDetector".isIpad
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isIpad = function() {
    return isDeviceType(
      config.vendors.apple.ipad.pattern.include,
      config.vendors.apple.ipad.pattern.exclude
    );
  };

  /**
   * Method to detect iPhone.
   * @function external:"jQuery.fn.deviceDetector".isIphone
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isIphone = function() {
    return isDeviceType(
      config.vendors.apple.iphone.pattern.include,
      config.vendors.apple.iphone.pattern.exclude
    );
  };

  /**
   * Method to detect iOS.
   * @function external:"jQuery.fn.deviceDetector".isIos
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isIos = function() {
    return isDeviceType(
      config.vendors.apple.ios.pattern.include,
      config.vendors.apple.ios.pattern.exclude
    );
  };

  /**
   * Method to detect Mac OS.
   * @function external:"jQuery.fn.deviceDetector".isMacos
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isMacos = function() {
    return isDeviceType(
      config.vendors.apple.macos.pattern.include,
      config.vendors.apple.macos.pattern.exclude
    );
  };

  /**
   * Method to detect Chrome.
   * @function external:"jQuery.fn.deviceDetector".isChrome
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isChrome = function() {
    return isDeviceType(
      config.vendors.google.chrome.pattern.include,
      config.vendors.google.chrome.pattern.exclude
    );
  };

  /**
   * Method to detect Android.
   * @function external:"jQuery.fn.deviceDetector".isAndroid
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isAndroid = function() {
    return isDeviceType(
      config.vendors.google.android.pattern.include,
      config.vendors.google.android.pattern.exclude
    );
  };

  /**
   * Method to detect Firefox.
   * @function external:"jQuery.fn.deviceDetector".isFirefox
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isFirefox = function() {
    return isDeviceType(
      config.vendors.mozilla.firefox.pattern.include,
      config.vendors.mozilla.firefox.pattern.exclude
    );
  };

  /**
   * Method to detect Microsoft Internet Explorer (IE/Edge).
   * @function external:"jQuery.fn.deviceDetector".isMsie
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isMsie = function() {
    return isDeviceType(
      config.vendors.microsoft.msie.pattern.include,
      config.vendors.microsoft.msie.pattern.exclude
    );
  };

  /**
   * Method to detect Microsoft Edge.
   * @function external:"jQuery.fn.deviceDetector".isEdge
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isEdge = function() {
    return isDeviceType(
      config.vendors.microsoft.edge.pattern.include,
      config.vendors.microsoft.edge.pattern.exclude
    );
  };

  /**
   * Method to detect Internet Explorer.
   * @function external:"jQuery.fn.deviceDetector".isIe
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isIe = function() {
    return isDeviceType(
      config.vendors.microsoft.ie.pattern.include,
      config.vendors.microsoft.ie.pattern.exclude
    );
  };

  /**
   * Method to detect Microsoft Internet Explorer Mobile.
   * @function external:"jQuery.fn.deviceDetector".isIeMobile
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isIeMobile = function() {
    return isDeviceType(
      config.vendors.microsoft.ieMobile.pattern.include,
      config.vendors.microsoft.ieMobile.pattern.exclude
    );
  };

  /**
   * Method to detect Windows.
   * @function external:"jQuery.fn.deviceDetector".isWindows
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isWindows = function() {
    return isDeviceType(
      config.vendors.microsoft.windows.pattern.include,
      config.vendors.microsoft.windows.pattern.exclude
    );
  };

  /**
   * Method to detect Windows Phone.
   * @function external:"jQuery.fn.deviceDetector".isWindowsPhone
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isWindowsPhone = function() {
    return isDeviceType(
      config.vendors.microsoft.windowsPhone.pattern.include,
      config.vendors.microsoft.windowsPhone.pattern.exclude
    );
  };

  /**
   * Method to detect Opera.
   * @function external:"jQuery.fn.deviceDetector".isOpera
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isOpera = function() {
    return isDeviceType(
      config.vendors.opera.opera.pattern.include,
      config.vendors.opera.opera.pattern.exclude
    );
  };

  /**
   * Method to detect Opera Mini.
   * @function external:"jQuery.fn.deviceDetector".isOperaMini
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isOperaMini = function() {
    return isDeviceType(
      config.vendors.opera.operaMini.pattern.include,
      config.vendors.opera.operaMini.pattern.exclude
    );
  };

  /**
   * Method to detect BlackBerry.
   * @function external:"jQuery.fn.deviceDetector".isBlackberry
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isBlackberry = function() {
    return isDeviceType(
      config.vendors.blackberry.blackberry.pattern.include,
      config.vendors.blackberry.blackberry.pattern.exclude
    );
  };

  /**
   * Method to detect Linux.
   * @function external:"jQuery.fn.deviceDetector".isLinux
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isLinux = function() {
    return isDeviceType(
      config.vendors.lbu.linux.pattern.include,
      config.vendors.lbu.linux.pattern.exclude
    );
  };

  /**
   * Method to detect BSD/Unix.
   * @function external:"jQuery.fn.deviceDetector".isBsd
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isBsd = function() {
    return isDeviceType(
      config.vendors.lbu.bsd.pattern.include,
      config.vendors.lbu.bsd.pattern.exclude
    );
  };

  /**
   * Method to get Browser Version.
   * @function external:"jQuery.fn.deviceDetector".getBrowserVersion
   * @return {string} The browser version.
   */

  $.fn.deviceDetector.getBrowserVersion = function() {
    return getBrowserVersion();
  };

  /**
   * Method to get Browser Name.
   * @function external:"jQuery.fn.deviceDetector".getBrowserName
   * @return {string} The browser name.
   */

  $.fn.deviceDetector.getBrowserName = function() {
    return getBrowserName();
  };

  /**
   * Method to get Browser Id (Short Name).
   * @function external:"jQuery.fn.deviceDetector".getBrowserId
   * @return {string} The browser id.
   */

  $.fn.deviceDetector.getBrowserId = function() {
    return getBrowserName(true);
  };

  /**
   * Method to get Operating System Version.
   * @function external:"jQuery.fn.deviceDetector".getOsVersion
   * @return {string} The OS version String.
   */

  $.fn.deviceDetector.getOsVersion = function() {
    return getOsVersion().string;
  };

  /**
   * Method to get Operating System Version String.
   * @function external:"jQuery.fn.deviceDetector".getOsVersionString
   * @return {string} The OS version String.
   */

  $.fn.deviceDetector.getOsVersionString = function() {
    return getOsVersion().string;
  };

  /**
   * Method to get Operating System Version Categories.
   * @function external:"jQuery.fn.deviceDetector".getOsVersionCategories
   * @return {object} The OS version Categories.
   */

  $.fn.deviceDetector.getOsVersionCategories = function() {
    return getOsVersion().categories;
  };

  /**
   * Method to get Operating System Version Major.
   * @function external:"jQuery.fn.deviceDetector".getOsVersionMajor
   * @return {number} The OS version Major.
   */

  $.fn.deviceDetector.getOsVersionMajor = function() {
    return getOsVersion().categories.major;
  };

  /**
   * Method to get Operating System Version Minor.
   * @function external:"jQuery.fn.deviceDetector".getOsVersionMinor
   * @return {number} The OS version Minor.
   */

  $.fn.deviceDetector.getOsVersionMinor = function() {
    return getOsVersion().categories.minor;
  };

  /**
   * Method to get Operating System Version Bugfix.
   * @function external:"jQuery.fn.deviceDetector".getOsVersionBugfix
   * @return {number} The OS version Bugfix.
   */

  $.fn.deviceDetector.getOsVersionBugfix = function() {
    return getOsVersion().categories.bugfix;
  };

  /**
   * Method to get Operating System Name.
   * @function external:"jQuery.fn.deviceDetector".getOsName
   * @return {string} The OS name.
   */

  $.fn.deviceDetector.getOsName = function() {
    return getOsName();
  };

  /**
   * Method to get Operating System Id (Short Name).
   * @function external:"jQuery.fn.deviceDetector".getOsId
   * @return {string} The OS id.
   */

  $.fn.deviceDetector.getOsId = function() {
    return getOsName(true);
  };

  /**
   * Method to detect Browser and Device support.
   * @function external:"jQuery.fn.deviceDetector".isSupported
   * @return {boolean} The match status.
   */

  $.fn.deviceDetector.isSupported = function() {
    return isSupported();
  };

  /**
   * Method to get all available device and browser informations.
   * @function external:"jQuery.fn.deviceDetector".getInfo
   * @return {object} The device and browser infromation.
   */

  $.fn.deviceDetector.getInfo = function() {
    return {
      'browserVersion': this.getBrowserVersion(),
      'browserName': this.getBrowserName(),
      'browserId': this.getBrowserId(),
      'osVersion': this.getOsVersion(),
      'osVersionString': this.getOsVersionString(),
      'osVersionCategories': this.getOsVersionCategories(),
      'osVersionMajor': this.getOsVersionMajor(),
      'osVersionMinor': this.getOsVersionMinor(),
      'osVersionBugfix': this.getOsVersionBugfix(),
      'osName': this.getOsName(),
      'osId': this.getOsId(),
      'supported': isSupported(),
      'mobile': this.isMobile(),
      'desktop': this.isDesktop(),
      'safari': this.isSafari(),
      'iphone': this.isIphone(),
      'ipad': this.isIpad(),
      'ios': this.isIos(),
      'macos': this.isMacos(),
      'chrome': this.isChrome(),
      'android': this.isAndroid(),
      'firefox': this.isFirefox(),
      'ie': this.isIe(),
      'msie': this.isMsie(),
      'edge': this.isEdge(),
      'ieMobile': this.isIeMobile(),
      'windowsPhone': this.isWindowsPhone(),
      'windows': this.isWindows(),
      'opera': this.isOpera(),
      'operaMini': this.isOperaMini(),
      'blackberry': this.isBlackberry(),
      'linux': this.isLinux(),
      'bsd': this.isBsd(),
    };
  };

  /**
   * Method to remove empty Values from a Array.
   * @function removeEmptyValuesFromArray
   * @param {array} arr The Array to remove empty Values.
   * @return {array} The new Array without empty Values.
   */
  function removeEmptyValuesFromArray(arr) {
    return $.map( arr, function(value) {
       return value === '' ? null : value;
   });
  }

  /* start: test-code */
  $.fn.deviceDetector.__removeEmptyValuesFromArray = removeEmptyValuesFromArray;
  /* end: test-code */

  /**
   * Method to detect Characters matching.
   * @function isMatching
   * @param {array} arr The Characters to match.
   * @return {boolean} The match status.
   */
  function isMatching(arr) {
   var status = false; // eslint-disable-line no-var
   var newArr = removeEmptyValuesFromArray(arr); // eslint-disable-line no-var
   $.each( newArr, function( index, value ) {
     value = ('' + value).toLowerCase();
     status =
       browserAgentString.indexOf(value) > -1 ||
       browserVersionString.indexOf(value) > -1;
     if (status === true) return false;
   });
   return status;
  }

  /* start: test-code */
   $.fn.deviceDetector.__isMatching = isMatching;
  /* end: test-code */

  /**
   * Method to detect Device Type.
   * @function isDeviceType
   * @param {array|string} includes The Device to match.
   * @param {array|string} excludes The Device should not match.
   * @return {boolean} The Device Type match status.
   */
  function isDeviceType(includes, excludes) {
    try {
      var hasIncludes = false; // eslint-disable-line no-var
      var hasExcludes = false; // eslint-disable-line no-var

      if (!$.isArray(includes)) includes = $.makeArray(includes);

      if (!$.isArray(excludes)) excludes = $.makeArray(excludes);

      hasExcludes = isMatching(excludes);

      if (hasExcludes === false) hasIncludes = isMatching(includes);

      return hasIncludes;
    } catch (error) {
      console.info( // eslint-disable-line no-console
        'deviceDetector: No match String || Array given in isDeviceType()',
        error
      );
    }
  }

  /* start: test-code */
  $.fn.deviceDetector.__isDeviceType = isDeviceType;
  /* end: test-code */

  /**
   * Method to get the Browser Version.
   * @function getBrowserVersion
   * @return {number} The browser Version. Default is 0.
   */
  function getBrowserVersion() {
    var version = 0; // eslint-disable-line no-var
    var data = // eslint-disable-line no-var
      browserAgentString +
      browserVersionString;
    $.each(config.browsers.versions, function(key, value) {
      var index = data.indexOf(value.index); // eslint-disable-line no-var
      if (index > -1) {
        version = parseFloat(
          data.substring(index + value.map).split('.')[0]
        );
        return false;
      }
    });
    return version;
  }

  /* start: test-code */
  $.fn.deviceDetector.__getBrowserVersion = getBrowserVersion;
  /* end: test-code */

  /**
   * Method to get the Browser Name.
   * @function getBrowserName
   * @param {boolean} returnId The Method should return an Id
   * (browser short name) instead of the Name (browser long name)
   * @return {string} The browser Name. Default ist unknown.
   */
  function getBrowserName(returnId) {
    var name = 'unknown'; // eslint-disable-line no-var
    $.each(config.browsers.names, function(key, value) {
      var isBrowser = isDeviceType( // eslint-disable-line no-var
        value.pattern.include,
        value.pattern.exclude
      );
      if (isBrowser) {
        (returnId === true) ? name = value.id : name = value.name;
        return false;
      }
    });
    return name;
  }

  /* start: test-code */
  $.fn.deviceDetector.__getBrowserName = getBrowserName;
  /* end: test-code */

  /**
   * Method to get the Operating System Version.
   * @function getOsVersion
   * @return {object} The OS Version . Default is 0.
   * {
   *  'string': '0',
   *  'categories': {
   *    'major': 0,
   *    'minor': 0,
   *    'bugfix': 0,
   *  }
   * }
   */
  function getOsVersion() {
    var version = { // eslint-disable-line no-var
      'string': '0.0.0',
      'categories': {
        'major': 0,
        'minor': 0,
        'bugfix': 0,
      },
    };
    var data = // eslint-disable-line no-var
      browserAgentString +
      browserVersionString;
    $.each(config.oss.versions, function(key, value) {
      var index = data.indexOf(value.index); // eslint-disable-line no-var
      if (index > -1) {
        version.string =
          data.substring(index + value.map).split(value.cut)[0]
          || '0.0.0';
        var choped = // eslint-disable-line no-var
          version.string.split(value.chop);
        version.categories.major = parseInt(choped[0]) || 0;
        version.categories.minor = parseInt(choped[1]) || 0;
        version.categories.bugfix = parseInt(choped[2]) || 0;
        return false;
      }
    });
    return version;
  }

  /* start: test-code */
  $.fn.deviceDetector.__getOsVersion = getOsVersion;
  /* end: test-code */

  /**
   * Method to get the Operating System Name.
   * @function getOsName
   * @param {boolean} returnId The Method should return an Id
   * (OS short name) instead of the Name (OS long name)
   * @return {string} The OS Name. Default ist unknown.
   */
  function getOsName(returnId) {
    var name = 'unknown'; // eslint-disable-line no-var
    $.each(config.oss.names, function(key, value) {
      var isOs = isDeviceType( // eslint-disable-line no-var
        value.pattern.include,
        value.pattern.exclude
      );
      if (isOs) {
        (returnId === true) ? name = value.id : name = value.name;
        return false;
      }
    });
    return name;
  }

  /* start: test-code */
  $.fn.deviceDetector.__getOsName = getOsName;
  /* end: test-code */

  /**
   * Method to detect supported Browser.
   * @function isSupported
   * @return {boolean} The supported Browser match status.
   */
  function isSupported() {
    var isSupported = false; // eslint-disable-line no-var
    $.each(config.supports, function(key, value) {
      if (
        getBrowserName(true) === value.id &&
        getBrowserVersion() >= parseFloat(value.version)
      ) isSupported = true;
    });
    return isSupported;
  }

  /* start: test-code */
  $.fn.deviceDetector.__isSupported = isSupported;
  /* end: test-code */

  // private variables
  var browser = navigator; // eslint-disable-line no-var
  var browserAgentString = // eslint-disable-line no-var
    ('' + browser.userAgent).toLowerCase();
  var browserVersionString = // eslint-disable-line no-var
    ('' + browser.appVersion).toLowerCase();
  browserAgentString = ('' + browserAgentString).toLowerCase();
  browserVersionString = browserAgentString || browserVersionString;

  // config
  //
  //  vendor
  //    Apple
  //    Google
  //    Mozilla
  //    Microsoft
  //    Opera
  //    Blackberry
  //    lbu (linux, bsd and unix)
  //  browsers
  //    names: browser nameings and patterns
  //    versions: index / matching patterns
  //  oss
  //    names: os nameings and patterns
  //    versions: index / matching patterns
  //  mobile
  //    mobile matching patterns
  //  supports
  //    supported browser / browser matrix
  var config = {}; // eslint-disable-line no-var
  $.extend(
    config,
    {
      'vendors': {
        'apple': {
          'safari': {
            'pattern': {
              'include': 'safari',
              'exclude': ['chrome', 'iemobile', 'opr/', 'opera'],
            },
          },
          'iphone': {
            'pattern': {
              'include': 'iphone',
              'exclude': 'iemobile',
            },
          },
          'ipad': {
            'pattern': {
              'include': 'ipad',
              'exclude': 'iemobile',
            },
          },
          'ios': {
            'pattern': {
              'include': ['ipad', 'iphone', 'ipod'],
              'exclude': 'iemobile',
            },
          },
          'macos': {
            'pattern': {
              'include': 'mac os',
              'exclude': ['iphone', 'ipad', 'ipod'],
            },
          },
        },
        'google': {
          'chrome': {
            'pattern': {
              'include': 'chrome',
              'exclude': ['edge', 'msie', 'firefox', 'opr/', 'opera'],
            },
          },
          'android': {
            'pattern': {
              'include': 'android',
              'exclude': 'windows phone',
            },
          },
        },
        'mozilla': {
          'firefox': {
            'pattern': {
              'include': 'firefox',
              'exclude': '',
            },
          },
        },
        'microsoft': {
          'msie': {
            'pattern': {
              'include': ['trident', 'msie'],
              'exclude': 'iemobile',
            },
          },
          'edge': {
            'pattern': {
              'include': 'edge',
              'exclude': 'iemobile',
            },
          },
          'ie': {
            'pattern': {
              'include': [
                'trident',
                'msie',
                'edge',
              ],
              'exclude': 'iemobile',
            },
          },
          'ieMobile': {
            'pattern': {
              'include': 'iemobile',
              'exclude': '',
            },
          },
          'windows': {
            'pattern': {
              'include': 'windows nt',
              'exclude': '',
            },
          },
          'windowsMobile': {
            'pattern': {
              'include': 'windows phone',
              'exclude': '',
            },
          },
          'windowsXp': {
            'pattern': {
              'include': 'windows nt 5',
              'exclude': '',
            },
          },
          'windowsVista': {
            'pattern': {
              'include': 'windows nt 6.0',
              'exclude': '',
            },
          },
          'windows7': {
            'pattern': {
              'include': 'windows nt 6.1',
              'exclude': '',
            },
          },
          'windows8': {
            'pattern': {
              'include': 'windows nt 6.2',
              'exclude': '',
            },
          },
          'windows81': {
            'pattern': {
              'include': 'windows nt 6.3',
              'exclude': '',
            },
          },
          'windows10': {
            'pattern': {
              'include': 'windows nt 10.0',
              'exclude': '',
            },
          },
          'windowsPhone': {
            'pattern': {
              'include': 'windows phone',
              'exclude': '',
            },
          },
        },
        'opera': {
          'opera': {
            'pattern': {
              'include': ['opera', 'opr', 'presto'],
              'exclude': 'opera mini',
            },
          },
          'operaMini': {
            'pattern': {
              'include': 'opera mini',
              'exclude': '',
            },
          },
        },
        'blackberry': {
          'blackberry': {
            'pattern': {
              'include': 'blackberry',
              'exclude': '',
            },
          },
        },
        'lbu': {
          'linux': {
            'pattern': {
              'include': 'linux',
              'exclude': '',
            },
          },
          'bsd': {
            'pattern': {
              'include': ['bsd', 'unix', 'sunos'],
              'exclude': '',
            },
          },
          'unixlike': {
            'pattern': {
              'include': ['linux', 'bsd', 'unix', 'sunos', 'X11'],
              'exclude': ['mobile', 'tablet', 'phone', 'droid'],
            },
          },
        },
      },
    }
  );
  $.extend(
    config,
    {
      'browsers': {
        'names': {
          'edge': {
            'id': 'edge',
            'name': 'Microsoft Edge',
            'pattern': config.vendors.microsoft.edge.pattern,
          },
          'ie': {
            'id': 'msie',
            'name': 'Microsoft Internet Explorer',
            'pattern': config.vendors.microsoft.ie.pattern,
          },
          'ieMobile': {
            'id': 'ieMobile',
            'name': 'Microsoft Internet Explorer Mobile',
            'pattern': config.vendors.microsoft.ieMobile.pattern,
          },
          'chrome': {
            'id': 'chrome',
            'name': 'Google Chrome',
            'pattern': config.vendors.google.chrome.pattern,
          },
          'safari': {
            'id': 'safari',
            'name': 'Apple Safari',
            'pattern': config.vendors.apple.safari.pattern,
          },
          'firefox': {
            'id': 'firefox',
            'name': 'Mozilla Firefox',
            'pattern': config.vendors.mozilla.firefox.pattern,
          },
          'opera': {
            'id': 'opera',
            'name': 'Opera',
            'pattern': config.vendors.opera.opera.pattern,
          },
          'operaMini': {
            'id': 'operaMini',
            'name': 'Opera Mini',
            'pattern': config.vendors.opera.operaMini.pattern,
          },
          'blackberry': {
            'id': 'blackberry',
            'name': 'BlackBerry',
            'pattern': config.vendors.blackberry.blackberry.pattern,
          },
        },
        'versions': {
          'default': {
            'index': 'rv:',
            'map': 3,
          },
          'edge': {
            'index': 'edge/',
            'map': 5,
          },
          'msie': {
            'index': 'msie ',
            'map': 5,
          },
          'ieMobile': {
            'index': 'iemobile/',
            'map': 9,
          },
          'chrome': {
            'index': 'chrome/',
            'map': 7,
          },
          'firefox': {
            'index': 'firefox/',
            'map': 8,
          },
          'opr': {
            'index': 'opr/',
            'map': 4,
          },
          'operaMini': {
            'index': 'opera/',
            'map': 6,
          },
          'opera': {
            'index': 'opera ',
            'map': 5,
          },
          'safari': {
            'index': 'version/',
            'map': 8,
          },
        },
      },
    }
  );
  $.extend(
    config,
    {
      'oss': {
        'names': {
          'windowsPhone': {
            'id': 'windowsPhone',
            'name': 'Microsoft Windows Phone',
            'pattern': config.vendors.microsoft.windowsPhone.pattern,
          },
          'windowsXp': {
            'id': 'windowsXp',
            'name': 'Microsoft Windows XP',
            'pattern': config.vendors.microsoft.windowsXp.pattern,
          },
          'windowsVista': {
            'id': 'windowsVista',
            'name': 'Microsoft Windows Vista',
            'pattern': config.vendors.microsoft.windowsVista.pattern,
          },
          'windows7': {
            'id': 'windows7',
            'name': 'Microsoft Windows 7',
            'pattern': config.vendors.microsoft.windows7.pattern,
          },
          'window8': {
            'id': 'windows8',
            'name': 'Microsoft Windows 8',
            'pattern': config.vendors.microsoft.windows8.pattern,
          },
          'window81': {
            'id': 'windows81',
            'name': 'Microsoft Windows 8.1',
            'pattern': config.vendors.microsoft.windows81.pattern,
          },
          'windows10': {
            'id': 'windows10',
            'name': 'Microsoft Windows 10',
            'pattern': config.vendors.microsoft.windows10.pattern,
          },
          'macos': {
            'id': 'macos',
            'name': 'Apple Mac OS X',
            'pattern': config.vendors.apple.macos.pattern,
          },
          'ios': {
            'id': 'ios',
            'name': 'Apple iOS',
            'pattern': config.vendors.apple.ios.pattern,
          },
          'android': {
            'id': 'android',
            'name': 'Android Linux',
            'pattern': config.vendors.google.android.pattern,
          },
          'linux': {
            'id': 'linux',
            'name': 'GNU/Linux OS',
            'pattern': config.vendors.lbu.linux.pattern,
          },
          'bsd': {
            'id': 'bsd',
            'name': 'BSD OS',
            'pattern': config.vendors.lbu.bsd.pattern,
          },
          'blackberry': {
            'id': 'blackberry',
            'name': 'BlackBerry OS',
            'pattern': config.vendors.blackberry.blackberry.pattern,
          },
        },
        'versions': {
          'windowsPhone': {
            'index': config.vendors.microsoft.windowsPhone.pattern.include,
            'map': 14,
            'cut': /;|\)/,
            'chop': '.',
          },
          'windowsXp': {
            'index': config.vendors.microsoft.windowsXp.pattern.include,
            'map': 11,
            'cut': /;|\)/,
            'chop': '.',
          },
          'windowsVista': {
            'index': config.vendors.microsoft.windowsVista.pattern.include,
            'map': 11,
            'cut': /;|\)/,
            'chop': '.',
          },
          'windows7': {
            'index': config.vendors.microsoft.windows7.pattern.include,
            'map': 11,
            'cut': /;|\)/,
            'chop': '.',
          },
          'windows8': {
            'index': config.vendors.microsoft.windows8.pattern.include,
            'map': 11,
            'cut': /;|\)/,
            'chop': '.',
          },
          'windows81': {
            'index': config.vendors.microsoft.windows81.pattern.include,
            'map': 11,
            'cut': /;|\)/,
            'chop': '.',
          },
          'windows10': {
            'index': config.vendors.microsoft.windows10.pattern.include,
            'map': 11,
            'cut': /;|\)/,
            'chop': '.',
          },
          'android': {
            'index': config.vendors.google.android.pattern.include,
            'map': 8,
            'cut': /;|\)/,
            'chop': '.',
          },
          'ios': {
            'index': 'cpu os ',
            'map': 7,
            'cut': / |\)/,
            'chop': '_',
          },
          'iphone': {
            'index': 'iphone os ',
            'map': 10,
            'cut': / |\)/,
            'chop': '_',
          },
          'ipad': {
            'index': 'ipad os ',
            'map': 8,
            'cut': / |\)/,
            'chop': '_',
          },
          'macos': {
            'index': 'mac os x',
            'map': 9,
            'cut': / |\)|;/,
            'chop': /_|\./,
          },
          'bsd': {
            'index': config.vendors.lbu.bsd.pattern.include,
            'map': 5,
            'cut': / |\)/,
            'chop': '.',
          },
          'linux': {
            'index': config.vendors.lbu.linux.pattern.include,
            'map': 5,
            'cut': /;| |\)/,
            'chop': '.',
          },
          'blackberry': {
            'index': config.vendors.blackberry.blackberry.pattern.include,
            'map': 10,
            'cut': /;|\)/,
            'chop': '.',
          },
        },
      },
    }
  );
  $.extend(
    config,
    {
      'mobile': {
        'pattern': {
          'include': ['mobile', 'tablet', 'phone', 'droid'],
          'exclude': '',
        },
      },
    }
  );
  $.extend(
    config,
    {
      'supports': {
        'msie': {'id': 'msie', 'version': '11'},
        'edge': {'id': 'edge', 'version': '12'},
        'chrome': {'id': 'chrome', 'version': '66'},
        'firefox': {'id': 'firefox', 'version': '60'},
        'safari': {'id': 'safari', 'version': '11'},
      },
    }
  );
}(jQuery));
