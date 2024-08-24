/*
 * derived from https://github.com/SteffTek/planet.js version 1.0.5
 * Apache-2.0 license
 */

!(function () {
  var t = {
      121: function (t, e, n) {
        (e.generatePerlinNoise = function (t, e, n) {
          var r,
            o = (n = n || {}).octaveCount || 4,
            i = n.amplitude || 0.1,
            u = n.persistence || 0.2,
            c = l(t, e),
            f = new Array(o);
          for (r = 0; r < o; ++r) f[r] = v(r);
          var s = new Array(t * e),
            d = 0;
          for (r = o - 1; r >= 0; --r) {
            d += i *= u;
            for (var h = 0; h < s.length; ++h)
              (s[h] = s[h] || 0), (s[h] += f[r][h] * i);
          }
          for (r = 0; r < s.length; ++r) s[r] /= d;
          return s;
          function v(n) {
            for (
              var r = new Array(t * e),
                o = Math.pow(2, n),
                l = 1 / o,
                i = 0,
                u = 0;
              u < e;
              ++u
            )
              for (
                var f = Math.floor(u / o) * o,
                  s = (f + o) % e,
                  d = (u - f) * l,
                  h = 0;
                h < t;
                ++h
              ) {
                var v = Math.floor(h / o) * o,
                  p = (v + o) % t,
                  g = (h - v) * l,
                  x = a(c[f * t + v], c[s * t + v], d),
                  m = a(c[f * t + p], c[s * t + p], d);
                (r[i] = a(x, m, g)), (i += 1);
              }
            return r;
          }
        }),
          (e.generateWhiteNoise = l),
          (e.setSeed = function (t) {
            o = t && null !== t ? r(t) : null;
          });
        const r = n(391);
        var o = null;
        function l(t, e) {
          for (var n = new Array(t * e), r = 0; r < n.length; ++r)
            n[r] = o && null !== o ? o() : Math.random();
          return n;
        }
        function a(t, e, n) {
          return t * (1 - n) + n * e;
        }
      },
      391: function (t, e, n) {
        var r = n(180),
          o = n(181),
          l = n(31),
          a = n(67),
          i = n(833),
          u = n(717),
          c = n(801);
        (c.alea = r),
          (c.xor128 = o),
          (c.xorwow = l),
          (c.xorshift7 = a),
          (c.xor4096 = i),
          (c.tychei = u),
          (t.exports = c);
      },
      180: function (t, e, n) {
        var r;
        !(function (t, o) {
          function l(t) {
            var e,
              n = this,
              r =
                ((e = 4022871197),
                function (t) {
                  t = t.toString();
                  for (var n = 0; n < t.length; n++) {
                    var r = 0.02519603282416938 * (e += t.charCodeAt(n));
                    (r -= e = r >>> 0),
                      (e = (r *= e) >>> 0),
                      (e += 4294967296 * (r -= e));
                  }
                  return 2.3283064365386963e-10 * (e >>> 0);
                });
            (n.next = function () {
              var t = 2091639 * n.s0 + 2.3283064365386963e-10 * n.c;
              return (n.s0 = n.s1), (n.s1 = n.s2), (n.s2 = t - (n.c = 0 | t));
            }),
              (n.c = 1),
              (n.s0 = r(" ")),
              (n.s1 = r(" ")),
              (n.s2 = r(" ")),
              (n.s0 -= r(t)),
              n.s0 < 0 && (n.s0 += 1),
              (n.s1 -= r(t)),
              n.s1 < 0 && (n.s1 += 1),
              (n.s2 -= r(t)),
              n.s2 < 0 && (n.s2 += 1),
              (r = null);
          }
          function a(t, e) {
            return (e.c = t.c), (e.s0 = t.s0), (e.s1 = t.s1), (e.s2 = t.s2), e;
          }
          function i(t, e) {
            var n = new l(t),
              r = e && e.state,
              o = n.next;
            return (
              (o.int32 = function () {
                return (4294967296 * n.next()) | 0;
              }),
              (o.double = function () {
                return o() + 11102230246251565e-32 * ((2097152 * o()) | 0);
              }),
              (o.quick = o),
              r &&
                ("object" == typeof r && a(r, n),
                (o.state = function () {
                  return a(n, {});
                })),
              o
            );
          }
          o && o.exports
            ? (o.exports = i)
            : n.amdD && n.amdO
              ? void 0 ===
                  (r = function () {
                    return i;
                  }.call(e, n, e, o)) || (o.exports = r)
              : (this.alea = i);
        })(0, (t = n.nmd(t)), n.amdD);
      },
      717: function (t, e, n) {
        var r;
        !(function (t, o) {
          function l(t) {
            var e = this,
              n = "";
            (e.next = function () {
              var t = e.b,
                n = e.c,
                r = e.d,
                o = e.a;
              return (
                (t = (t << 25) ^ (t >>> 7) ^ n),
                (n = (n - r) | 0),
                (r = (r << 24) ^ (r >>> 8) ^ o),
                (o = (o - t) | 0),
                (e.b = t = (t << 20) ^ (t >>> 12) ^ n),
                (e.c = n = (n - r) | 0),
                (e.d = (r << 16) ^ (n >>> 16) ^ o),
                (e.a = (o - t) | 0)
              );
            }),
              (e.a = 0),
              (e.b = 0),
              (e.c = -1640531527),
              (e.d = 1367130551),
              t === Math.floor(t)
                ? ((e.a = (t / 4294967296) | 0), (e.b = 0 | t))
                : (n += t);
            for (var r = 0; r < n.length + 20; r++)
              (e.b ^= 0 | n.charCodeAt(r)), e.next();
          }
          function a(t, e) {
            return (e.a = t.a), (e.b = t.b), (e.c = t.c), (e.d = t.d), e;
          }
          function i(t, e) {
            var n = new l(t),
              r = e && e.state,
              o = function () {
                return (n.next() >>> 0) / 4294967296;
              };
            return (
              (o.double = function () {
                do {
                  var t =
                    ((n.next() >>> 11) + (n.next() >>> 0) / 4294967296) /
                    (1 << 21);
                } while (0 === t);
                return t;
              }),
              (o.int32 = n.next),
              (o.quick = o),
              r &&
                ("object" == typeof r && a(r, n),
                (o.state = function () {
                  return a(n, {});
                })),
              o
            );
          }
          o && o.exports
            ? (o.exports = i)
            : n.amdD && n.amdO
              ? void 0 ===
                  (r = function () {
                    return i;
                  }.call(e, n, e, o)) || (o.exports = r)
              : (this.tychei = i);
        })(0, (t = n.nmd(t)), n.amdD);
      },
      181: function (t, e, n) {
        var r;
        !(function (t, o) {
          function l(t) {
            var e = this,
              n = "";
            (e.x = 0),
              (e.y = 0),
              (e.z = 0),
              (e.w = 0),
              (e.next = function () {
                var t = e.x ^ (e.x << 11);
                return (
                  (e.x = e.y),
                  (e.y = e.z),
                  (e.z = e.w),
                  (e.w ^= (e.w >>> 19) ^ t ^ (t >>> 8))
                );
              }),
              t === (0 | t) ? (e.x = t) : (n += t);
            for (var r = 0; r < n.length + 64; r++)
              (e.x ^= 0 | n.charCodeAt(r)), e.next();
          }
          function a(t, e) {
            return (e.x = t.x), (e.y = t.y), (e.z = t.z), (e.w = t.w), e;
          }
          function i(t, e) {
            var n = new l(t),
              r = e && e.state,
              o = function () {
                return (n.next() >>> 0) / 4294967296;
              };
            return (
              (o.double = function () {
                do {
                  var t =
                    ((n.next() >>> 11) + (n.next() >>> 0) / 4294967296) /
                    (1 << 21);
                } while (0 === t);
                return t;
              }),
              (o.int32 = n.next),
              (o.quick = o),
              r &&
                ("object" == typeof r && a(r, n),
                (o.state = function () {
                  return a(n, {});
                })),
              o
            );
          }
          o && o.exports
            ? (o.exports = i)
            : n.amdD && n.amdO
              ? void 0 ===
                  (r = function () {
                    return i;
                  }.call(e, n, e, o)) || (o.exports = r)
              : (this.xor128 = i);
        })(0, (t = n.nmd(t)), n.amdD);
      },
      833: function (t, e, n) {
        var r;
        !(function (t, o) {
          function l(t) {
            var e = this;
            (e.next = function () {
              var t,
                n,
                r = e.w,
                o = e.X,
                l = e.i;
              return (
                (e.w = r = (r + 1640531527) | 0),
                (n = o[(l + 34) & 127]),
                (t = o[(l = (l + 1) & 127)]),
                (n ^= n << 13),
                (t ^= t << 17),
                (n ^= n >>> 15),
                (t ^= t >>> 12),
                (n = o[l] = n ^ t),
                (e.i = l),
                (n + (r ^ (r >>> 16))) | 0
              );
            }),
              (function (t, e) {
                var n,
                  r,
                  o,
                  l,
                  a,
                  i = [],
                  u = 128;
                for (
                  e === (0 | e)
                    ? ((r = e), (e = null))
                    : ((e += "\0"), (r = 0), (u = Math.max(u, e.length))),
                    o = 0,
                    l = -32;
                  l < u;
                  ++l
                )
                  e && (r ^= e.charCodeAt((l + 32) % e.length)),
                    0 === l && (a = r),
                    (r ^= r << 10),
                    (r ^= r >>> 15),
                    (r ^= r << 4),
                    (r ^= r >>> 13),
                    l >= 0 &&
                      ((a = (a + 1640531527) | 0),
                      (o = 0 == (n = i[127 & l] ^= r + a) ? o + 1 : 0));
                for (
                  o >= 128 && (i[127 & ((e && e.length) || 0)] = -1),
                    o = 127,
                    l = 512;
                  l > 0;
                  --l
                )
                  (r = i[(o + 34) & 127]),
                    (n = i[(o = (o + 1) & 127)]),
                    (r ^= r << 13),
                    (n ^= n << 17),
                    (r ^= r >>> 15),
                    (n ^= n >>> 12),
                    (i[o] = r ^ n);
                (t.w = a), (t.X = i), (t.i = o);
              })(e, t);
          }
          function a(t, e) {
            return (e.i = t.i), (e.w = t.w), (e.X = t.X.slice()), e;
          }
          function i(t, e) {
            null == t && (t = +new Date());
            var n = new l(t),
              r = e && e.state,
              o = function () {
                return (n.next() >>> 0) / 4294967296;
              };
            return (
              (o.double = function () {
                do {
                  var t =
                    ((n.next() >>> 11) + (n.next() >>> 0) / 4294967296) /
                    (1 << 21);
                } while (0 === t);
                return t;
              }),
              (o.int32 = n.next),
              (o.quick = o),
              r &&
                (r.X && a(r, n),
                (o.state = function () {
                  return a(n, {});
                })),
              o
            );
          }
          o && o.exports
            ? (o.exports = i)
            : n.amdD && n.amdO
              ? void 0 ===
                  (r = function () {
                    return i;
                  }.call(e, n, e, o)) || (o.exports = r)
              : (this.xor4096 = i);
        })(0, (t = n.nmd(t)), n.amdD);
      },
      67: function (t, e, n) {
        var r;
        !(function (t, o) {
          function l(t) {
            var e = this;
            (e.next = function () {
              var t,
                n,
                r = e.x,
                o = e.i;
              return (
                (t = r[o]),
                (n = (t ^= t >>> 7) ^ (t << 24)),
                (n ^= (t = r[(o + 1) & 7]) ^ (t >>> 10)),
                (n ^= (t = r[(o + 3) & 7]) ^ (t >>> 3)),
                (n ^= (t = r[(o + 4) & 7]) ^ (t << 7)),
                (t = r[(o + 7) & 7]),
                (n ^= (t ^= t << 13) ^ (t << 9)),
                (r[o] = n),
                (e.i = (o + 1) & 7),
                n
              );
            }),
              (function (t, e) {
                var n,
                  r = [];
                if (e === (0 | e)) r[0] = e;
                else
                  for (e = "" + e, n = 0; n < e.length; ++n)
                    r[7 & n] =
                      (r[7 & n] << 15) ^
                      ((e.charCodeAt(n) + r[(n + 1) & 7]) << 13);
                for (; r.length < 8; ) r.push(0);
                for (n = 0; n < 8 && 0 === r[n]; ++n);
                for (
                  8 == n ? (r[7] = -1) : r[n], t.x = r, t.i = 0, n = 256;
                  n > 0;
                  --n
                )
                  t.next();
              })(e, t);
          }
          function a(t, e) {
            return (e.x = t.x.slice()), (e.i = t.i), e;
          }
          function i(t, e) {
            null == t && (t = +new Date());
            var n = new l(t),
              r = e && e.state,
              o = function () {
                return (n.next() >>> 0) / 4294967296;
              };
            return (
              (o.double = function () {
                do {
                  var t =
                    ((n.next() >>> 11) + (n.next() >>> 0) / 4294967296) /
                    (1 << 21);
                } while (0 === t);
                return t;
              }),
              (o.int32 = n.next),
              (o.quick = o),
              r &&
                (r.x && a(r, n),
                (o.state = function () {
                  return a(n, {});
                })),
              o
            );
          }
          o && o.exports
            ? (o.exports = i)
            : n.amdD && n.amdO
              ? void 0 ===
                  (r = function () {
                    return i;
                  }.call(e, n, e, o)) || (o.exports = r)
              : (this.xorshift7 = i);
        })(0, (t = n.nmd(t)), n.amdD);
      },
      31: function (t, e, n) {
        var r;
        !(function (t, o) {
          function l(t) {
            var e = this,
              n = "";
            (e.next = function () {
              var t = e.x ^ (e.x >>> 2);
              return (
                (e.x = e.y),
                (e.y = e.z),
                (e.z = e.w),
                (e.w = e.v),
                ((e.d = (e.d + 362437) | 0) +
                  (e.v = e.v ^ (e.v << 4) ^ t ^ (t << 1))) |
                  0
              );
            }),
              (e.x = 0),
              (e.y = 0),
              (e.z = 0),
              (e.w = 0),
              (e.v = 0),
              t === (0 | t) ? (e.x = t) : (n += t);
            for (var r = 0; r < n.length + 64; r++)
              (e.x ^= 0 | n.charCodeAt(r)),
                r == n.length && (e.d = (e.x << 10) ^ (e.x >>> 4)),
                e.next();
          }
          function a(t, e) {
            return (
              (e.x = t.x),
              (e.y = t.y),
              (e.z = t.z),
              (e.w = t.w),
              (e.v = t.v),
              (e.d = t.d),
              e
            );
          }
          function i(t, e) {
            var n = new l(t),
              r = e && e.state,
              o = function () {
                return (n.next() >>> 0) / 4294967296;
              };
            return (
              (o.double = function () {
                do {
                  var t =
                    ((n.next() >>> 11) + (n.next() >>> 0) / 4294967296) /
                    (1 << 21);
                } while (0 === t);
                return t;
              }),
              (o.int32 = n.next),
              (o.quick = o),
              r &&
                ("object" == typeof r && a(r, n),
                (o.state = function () {
                  return a(n, {});
                })),
              o
            );
          }
          o && o.exports
            ? (o.exports = i)
            : n.amdD && n.amdO
              ? void 0 ===
                  (r = function () {
                    return i;
                  }.call(e, n, e, o)) || (o.exports = r)
              : (this.xorwow = i);
        })(0, (t = n.nmd(t)), n.amdD);
      },
      801: function (t, e, n) {
        var r;
        !(function (o, l) {
          var a,
            i = (0, eval)("this"),
            u = 256,
            c = "random",
            f = l.pow(u, 6),
            s = l.pow(2, 52),
            d = 2 * s,
            h = u - 1;
          function v(t, e, n) {
            var r = [],
              h = m(
                x(
                  (e = 1 == e ? { entropy: !0 } : e || {}).entropy
                    ? [t, _(o)]
                    : null == t
                      ? (function () {
                          try {
                            var t;
                            return (
                              a && (t = a.randomBytes)
                                ? (t = t(u))
                                : ((t = new Uint8Array(u)),
                                  (i.crypto || i.msCrypto).getRandomValues(t)),
                              _(t)
                            );
                          } catch (t) {
                            var e = i.navigator,
                              n = e && e.plugins;
                            return [+new Date(), i, n, i.screen, _(o)];
                          }
                        })()
                      : t,
                  3,
                ),
                r,
              ),
              v = new p(r),
              w = function () {
                for (var t = v.g(6), e = f, n = 0; t < s; )
                  (t = (t + n) * u), (e *= u), (n = v.g(1));
                for (; t >= d; ) (t /= 2), (e /= 2), (n >>>= 1);
                return (t + n) / e;
              };
            return (
              (w.int32 = function () {
                return 0 | v.g(4);
              }),
              (w.quick = function () {
                return v.g(4) / 4294967296;
              }),
              (w.double = w),
              m(_(v.S), o),
              (
                e.pass ||
                n ||
                function (t, e, n, r) {
                  return (
                    r &&
                      (r.S && g(r, v),
                      (t.state = function () {
                        return g(v, {});
                      })),
                    n ? ((l[c] = t), e) : t
                  );
                }
              )(w, h, "global" in e ? e.global : this == l, e.state)
            );
          }
          function p(t) {
            var e,
              n = t.length,
              r = this,
              o = 0,
              l = (r.i = r.j = 0),
              a = (r.S = []);
            for (n || (t = [n++]); o < u; ) a[o] = o++;
            for (o = 0; o < u; o++)
              (a[o] = a[(l = h & (l + t[o % n] + (e = a[o])))]), (a[l] = e);
            (r.g = function (t) {
              for (var e, n = 0, o = r.i, l = r.j, a = r.S; t--; )
                (e = a[(o = h & (o + 1))]),
                  (n =
                    n * u +
                    a[h & ((a[o] = a[(l = h & (l + e))]) + (a[l] = e))]);
              return (r.i = o), (r.j = l), n;
            })(u);
          }
          function g(t, e) {
            return (e.i = t.i), (e.j = t.j), (e.S = t.S.slice()), e;
          }
          function x(t, e) {
            var n,
              r = [],
              o = typeof t;
            if (e && "object" == o)
              for (n in t)
                try {
                  r.push(x(t[n], e - 1));
                } catch (t) {}
            return r.length ? r : "string" == o ? t : t + "\0";
          }
          function m(t, e) {
            for (var n, r = t + "", o = 0; o < r.length; )
              e[h & o] = h & ((n ^= 19 * e[h & o]) + r.charCodeAt(o++));
            return _(e);
          }
          function _(t) {
            return String.fromCharCode.apply(0, t);
          }
          if (((l["seed" + c] = v), m(l.random(), o), t.exports)) {
            t.exports = v;
            try {
              a = n(234);
            } catch (t) {}
          } else
            void 0 ===
              (r = function () {
                return v;
              }.call(e, n, e, t)) || (t.exports = r);
        })([], Math);
      },
      823: function (t, e, n) {
        var r;
        !(function () {
          "use strict";
          var o = 0.5 * (Math.sqrt(3) - 1),
            l = (3 - Math.sqrt(3)) / 6,
            a = 1 / 3,
            i = 1 / 6,
            u = (Math.sqrt(5) - 1) / 4,
            c = (5 - Math.sqrt(5)) / 20;
          function f(t) {
            var e;
            (e =
              "function" == typeof t
                ? t
                : t
                  ? (function () {
                      var t,
                        e = 0,
                        n = 0,
                        r = 0,
                        o = 1,
                        l =
                          ((t = 4022871197),
                          function (e) {
                            e = e.toString();
                            for (var n = 0; n < e.length; n++) {
                              var r =
                                0.02519603282416938 * (t += e.charCodeAt(n));
                              (r -= t = r >>> 0),
                                (t = (r *= t) >>> 0),
                                (t += 4294967296 * (r -= t));
                            }
                            return 2.3283064365386963e-10 * (t >>> 0);
                          });
                      (e = l(" ")), (n = l(" ")), (r = l(" "));
                      for (var a = 0; a < arguments.length; a++)
                        (e -= l(arguments[a])) < 0 && (e += 1),
                          (n -= l(arguments[a])) < 0 && (n += 1),
                          (r -= l(arguments[a])) < 0 && (r += 1);
                      return (
                        (l = null),
                        function () {
                          var t = 2091639 * e + 2.3283064365386963e-10 * o;
                          return (e = n), (n = r), (r = t - (o = 0 | t));
                        }
                      );
                    })(t)
                  : Math.random),
              (this.p = s(e)),
              (this.perm = new Uint8Array(512)),
              (this.permMod12 = new Uint8Array(512));
            for (var n = 0; n < 512; n++)
              (this.perm[n] = this.p[255 & n]),
                (this.permMod12[n] = this.perm[n] % 12);
          }
          function s(t) {
            var e,
              n = new Uint8Array(256);
            for (e = 0; e < 256; e++) n[e] = e;
            for (e = 0; e < 255; e++) {
              var r = e + ~~(t() * (256 - e)),
                o = n[e];
              (n[e] = n[r]), (n[r] = o);
            }
            return n;
          }
          (f.prototype = {
            grad3: new Float32Array([
              1, 1, 0, -1, 1, 0, 1, -1, 0, -1, -1, 0, 1, 0, 1, -1, 0, 1, 1, 0,
              -1, -1, 0, -1, 0, 1, 1, 0, -1, 1, 0, 1, -1, 0, -1, -1,
            ]),
            grad4: new Float32Array([
              0, 1, 1, 1, 0, 1, 1, -1, 0, 1, -1, 1, 0, 1, -1, -1, 0, -1, 1, 1,
              0, -1, 1, -1, 0, -1, -1, 1, 0, -1, -1, -1, 1, 0, 1, 1, 1, 0, 1,
              -1, 1, 0, -1, 1, 1, 0, -1, -1, -1, 0, 1, 1, -1, 0, 1, -1, -1, 0,
              -1, 1, -1, 0, -1, -1, 1, 1, 0, 1, 1, 1, 0, -1, 1, -1, 0, 1, 1, -1,
              0, -1, -1, 1, 0, 1, -1, 1, 0, -1, -1, -1, 0, 1, -1, -1, 0, -1, 1,
              1, 1, 0, 1, 1, -1, 0, 1, -1, 1, 0, 1, -1, -1, 0, -1, 1, 1, 0, -1,
              1, -1, 0, -1, -1, 1, 0, -1, -1, -1, 0,
            ]),
            noise2D: function (t, e) {
              var n,
                r,
                a = this.permMod12,
                i = this.perm,
                u = this.grad3,
                c = 0,
                f = 0,
                s = 0,
                d = (t + e) * o,
                h = Math.floor(t + d),
                v = Math.floor(e + d),
                p = (h + v) * l,
                g = t - (h - p),
                x = e - (v - p);
              g > x ? ((n = 1), (r = 0)) : ((n = 0), (r = 1));
              var m = g - n + l,
                _ = x - r + l,
                w = g - 1 + 2 * l,
                y = x - 1 + 2 * l,
                b = 255 & h,
                S = 255 & v,
                M = 0.5 - g * g - x * x;
              if (M >= 0) {
                var k = 3 * a[b + i[S]];
                c = (M *= M) * M * (u[k] * g + u[k + 1] * x);
              }
              var C = 0.5 - m * m - _ * _;
              if (C >= 0) {
                var j = 3 * a[b + n + i[S + r]];
                f = (C *= C) * C * (u[j] * m + u[j + 1] * _);
              }
              var D = 0.5 - w * w - y * y;
              if (D >= 0) {
                var P = 3 * a[b + 1 + i[S + 1]];
                s = (D *= D) * D * (u[P] * w + u[P + 1] * y);
              }
              return 70 * (c + f + s);
            },
            noise3D: function (t, e, n) {
              var r,
                o,
                l,
                u,
                c,
                f,
                s,
                d,
                h,
                v,
                p = this.permMod12,
                g = this.perm,
                x = this.grad3,
                m = (t + e + n) * a,
                _ = Math.floor(t + m),
                w = Math.floor(e + m),
                y = Math.floor(n + m),
                b = (_ + w + y) * i,
                S = t - (_ - b),
                M = e - (w - b),
                k = n - (y - b);
              S >= M
                ? M >= k
                  ? ((c = 1), (f = 0), (s = 0), (d = 1), (h = 1), (v = 0))
                  : S >= k
                    ? ((c = 1), (f = 0), (s = 0), (d = 1), (h = 0), (v = 1))
                    : ((c = 0), (f = 0), (s = 1), (d = 1), (h = 0), (v = 1))
                : M < k
                  ? ((c = 0), (f = 0), (s = 1), (d = 0), (h = 1), (v = 1))
                  : S < k
                    ? ((c = 0), (f = 1), (s = 0), (d = 0), (h = 1), (v = 1))
                    : ((c = 0), (f = 1), (s = 0), (d = 1), (h = 1), (v = 0));
              var C = S - c + i,
                j = M - f + i,
                D = k - s + i,
                P = S - d + 2 * i,
                A = M - h + 2 * i,
                O = k - v + 2 * i,
                G = S - 1 + 0.5,
                I = M - 1 + 0.5,
                T = k - 1 + 0.5,
                N = 255 & _,
                q = 255 & w,
                F = 255 & y,
                R = 0.6 - S * S - M * M - k * k;
              if (R < 0) r = 0;
              else {
                var z = 3 * p[N + g[q + g[F]]];
                r = (R *= R) * R * (x[z] * S + x[z + 1] * M + x[z + 2] * k);
              }
              var B = 0.6 - C * C - j * j - D * D;
              if (B < 0) o = 0;
              else {
                var E = 3 * p[N + c + g[q + f + g[F + s]]];
                o = (B *= B) * B * (x[E] * C + x[E + 1] * j + x[E + 2] * D);
              }
              var U = 0.6 - P * P - A * A - O * O;
              if (U < 0) l = 0;
              else {
                var W = 3 * p[N + d + g[q + h + g[F + v]]];
                l = (U *= U) * U * (x[W] * P + x[W + 1] * A + x[W + 2] * O);
              }
              var X = 0.6 - G * G - I * I - T * T;
              if (X < 0) u = 0;
              else {
                var J = 3 * p[N + 1 + g[q + 1 + g[F + 1]]];
                u = (X *= X) * X * (x[J] * G + x[J + 1] * I + x[J + 2] * T);
              }
              return 32 * (r + o + l + u);
            },
            noise4D: function (t, e, n, r) {
              var o,
                l,
                a,
                i,
                f,
                s,
                d,
                h,
                v,
                p,
                g,
                x,
                m,
                _,
                w,
                y,
                b,
                S = this.perm,
                M = this.grad4,
                k = (t + e + n + r) * u,
                C = Math.floor(t + k),
                j = Math.floor(e + k),
                D = Math.floor(n + k),
                P = Math.floor(r + k),
                A = (C + j + D + P) * c,
                O = t - (C - A),
                G = e - (j - A),
                I = n - (D - A),
                T = r - (P - A),
                N = 0,
                q = 0,
                F = 0,
                R = 0;
              O > G ? N++ : q++,
                O > I ? N++ : F++,
                O > T ? N++ : R++,
                G > I ? q++ : F++,
                G > T ? q++ : R++,
                I > T ? F++ : R++;
              var z = O - (s = N >= 3 ? 1 : 0) + c,
                B = G - (d = q >= 3 ? 1 : 0) + c,
                E = I - (h = F >= 3 ? 1 : 0) + c,
                U = T - (v = R >= 3 ? 1 : 0) + c,
                W = O - (p = N >= 2 ? 1 : 0) + 2 * c,
                X = G - (g = q >= 2 ? 1 : 0) + 2 * c,
                J = I - (x = F >= 2 ? 1 : 0) + 2 * c,
                V = T - (m = R >= 2 ? 1 : 0) + 2 * c,
                H = O - (_ = N >= 1 ? 1 : 0) + 3 * c,
                K = G - (w = q >= 1 ? 1 : 0) + 3 * c,
                L = I - (y = F >= 1 ? 1 : 0) + 3 * c,
                Q = T - (b = R >= 1 ? 1 : 0) + 3 * c,
                Y = O - 1 + 4 * c,
                Z = G - 1 + 4 * c,
                $ = I - 1 + 4 * c,
                tt = T - 1 + 4 * c,
                et = 255 & C,
                nt = 255 & j,
                rt = 255 & D,
                ot = 255 & P,
                lt = 0.6 - O * O - G * G - I * I - T * T;
              if (lt < 0) o = 0;
              else {
                var at = (S[et + S[nt + S[rt + S[ot]]]] % 32) * 4;
                o =
                  (lt *= lt) *
                  lt *
                  (M[at] * O + M[at + 1] * G + M[at + 2] * I + M[at + 3] * T);
              }
              var it = 0.6 - z * z - B * B - E * E - U * U;
              if (it < 0) l = 0;
              else {
                var ut =
                  (S[et + s + S[nt + d + S[rt + h + S[ot + v]]]] % 32) * 4;
                l =
                  (it *= it) *
                  it *
                  (M[ut] * z + M[ut + 1] * B + M[ut + 2] * E + M[ut + 3] * U);
              }
              var ct = 0.6 - W * W - X * X - J * J - V * V;
              if (ct < 0) a = 0;
              else {
                var ft =
                  (S[et + p + S[nt + g + S[rt + x + S[ot + m]]]] % 32) * 4;
                a =
                  (ct *= ct) *
                  ct *
                  (M[ft] * W + M[ft + 1] * X + M[ft + 2] * J + M[ft + 3] * V);
              }
              var st = 0.6 - H * H - K * K - L * L - Q * Q;
              if (st < 0) i = 0;
              else {
                var dt =
                  (S[et + _ + S[nt + w + S[rt + y + S[ot + b]]]] % 32) * 4;
                i =
                  (st *= st) *
                  st *
                  (M[dt] * H + M[dt + 1] * K + M[dt + 2] * L + M[dt + 3] * Q);
              }
              var ht = 0.6 - Y * Y - Z * Z - $ * $ - tt * tt;
              if (ht < 0) f = 0;
              else {
                var vt =
                  (S[et + 1 + S[nt + 1 + S[rt + 1 + S[ot + 1]]]] % 32) * 4;
                f =
                  (ht *= ht) *
                  ht *
                  (M[vt] * Y + M[vt + 1] * Z + M[vt + 2] * $ + M[vt + 3] * tt);
              }
              return 27 * (o + l + a + i + f);
            },
          }),
            (f._buildPermutationTable = s),
            void 0 ===
              (r = function () {
                return f;
              }.call(e, n, e, t)) || (t.exports = r),
            (e.SimplexNoise = f),
            (t.exports = f);
        })();
      },
      44: function (t, e, n) {
        "use strict";
        t = n.hmd(t);
        const r = n(823),
          o = n(121),
          l = n(462),
          a = n(603);
        function i(t, e, n, r = null) {
          let l = [];
          o.setSeed(r);
          let a = o.generatePerlinNoise(t, e, n);
          for (let n = 0; n < e; n++) {
            let e = [];
            for (let r = 0; r < t; r++) e.push(a[n * t + r]);
            l.push(e);
          }
          return l;
        }
        function u(t, e, n) {
          let o,
            l,
            a = new r(n),
            i = [];
          for (l = 0; l < e; l++) {
            let e = [];
            for (o = 0; o < t; o++) e.push(a.noise2D(o, l));
            i.push(e);
          }
          return i;
        }
        function c(t, e, n, r) {
          let o = i(t, e, n, r),
            l = [];
          for (let e = 0; e < o.length; e++) {
            let n = f(o[e]);
            l.push([e, n]), (e += Math.round(t / 10));
          }
          return l;
        }
        function f(t) {
          if (0 === t.length) return -1;
          for (var e = t[0], n = 0, r = 1; r < t.length; r++)
            t[r] > e && ((n = r), (e = t[r]));
          return n;
        }
        function s(t, e) {
          let n = parseInt(t.replace("#", ""), 16),
            r = Math.round(2.55 * e),
            o = (n >> 16) + r,
            l = ((n >> 8) & 255) + r,
            a = (255 & n) + r;
          return (
            "#" +
            (
              16777216 +
              65536 * (o < 255 ? (o < 1 ? 0 : o) : 255) +
              256 * (l < 255 ? (l < 1 ? 0 : l) : 255) +
              (a < 255 ? (a < 1 ? 0 : a) : 255)
            )
              .toString(16)
              .slice(1)
          );
        }
        function d(t) {
          return (
            t < 0 && (t = 0),
            t > 100 && (t = 100),
            (t /= 100),
            1 == (t = (t = (t *= 255).toString(16)).split(".")[0]).length &&
              (t = "0" + t),
            t
          );
        }
        function h(t) {
          let e = parseInt(t.replace("#", ""), 16);
          return [e >> 16, 255 & e, (e >> 8) & 255];
        }
        function v(t, e, n) {
          return (
            "#" +
            (
              16777216 +
              65536 * (t < 255 ? (t < 1 ? 0 : t) : 255) +
              256 * (n < 255 ? (n < 1 ? 0 : n) : 255) +
              (e < 255 ? (e < 1 ? 0 : e) : 255)
            )
              .toString(16)
              .slice(1)
          );
        }
        function p(t, e, n) {
          let r = h(t),
            o = h(e),
            l = [
              Number.parseInt(r[0] * (1 - n) + o[0] * n),
              Number.parseInt(r[1] * (1 - n) + o[1] * n),
              Number.parseInt(r[2] * (1 - n) + o[2] * n),
            ];
          return v(l[0], l[1], l[2]);
        }
        function g(t, e, n, r) {
          let o,
            l,
            a,
            i,
            u,
            c,
            f,
            s = parseInt,
            d = Math.round,
            h = "string" == typeof n;
          return "number" != typeof t ||
            t < -1 ||
            t > 1 ||
            "string" != typeof e ||
            ("r" != e[0] && "#" != e[0]) ||
            (n && !h)
            ? null
            : (this.pSBCr ||
                (this.pSBCr = (t) => {
                  let e = t.length,
                    n = {};
                  if (e > 9) {
                    if (
                      (([o, l, a, h] = t = t.split(",")),
                      (e = t.length),
                      e < 3 || e > 4)
                    )
                      return null;
                    (n.r = s("a" == o[3] ? o.slice(5) : o.slice(4))),
                      (n.g = s(l)),
                      (n.b = s(a)),
                      (n.a = h ? parseFloat(h) : -1);
                  } else {
                    if (8 == e || 6 == e || e < 4) return null;
                    e < 6 &&
                      (t =
                        "#" +
                        t[1] +
                        t[1] +
                        t[2] +
                        t[2] +
                        t[3] +
                        t[3] +
                        (e > 4 ? t[4] + t[4] : "")),
                      (t = s(t.slice(1), 16)),
                      9 == e || 5 == e
                        ? ((n.r = (t >> 24) & 255),
                          (n.g = (t >> 16) & 255),
                          (n.b = (t >> 8) & 255),
                          (n.a = d((255 & t) / 0.255) / 1e3))
                        : ((n.r = t >> 16),
                          (n.g = (t >> 8) & 255),
                          (n.b = 255 & t),
                          (n.a = -1));
                  }
                  return n;
                }),
              (f = e.length > 9),
              (f = h ? n.length > 9 || ("c" == n && !f) : f),
              (u = this.pSBCr(e)),
              (i = t < 0),
              (c =
                n && "c" != n
                  ? this.pSBCr(n)
                  : i
                    ? { r: 0, g: 0, b: 0, a: -1 }
                    : { r: 255, g: 255, b: 255, a: -1 }),
              (i = 1 - (t = i ? -1 * t : t)),
              u && c
                ? (r
                    ? ((o = d(i * u.r + t * c.r)),
                      (l = d(i * u.g + t * c.g)),
                      (a = d(i * u.b + t * c.b)))
                    : ((o = d((i * u.r ** 2 + t * c.r ** 2) ** 0.5)),
                      (l = d((i * u.g ** 2 + t * c.g ** 2) ** 0.5)),
                      (a = d((i * u.b ** 2 + t * c.b ** 2) ** 0.5))),
                  (h = u.a),
                  (c = c.a),
                  (u = h >= 0 || c >= 0),
                  (h = u ? (h < 0 ? c : c < 0 ? h : h * i + c * t) : 0),
                  f
                    ? "rgb" +
                      (u ? "a(" : "(") +
                      o +
                      "," +
                      l +
                      "," +
                      a +
                      (u ? "," + d(1e3 * h) / 1e3 : "") +
                      ")"
                    : "#" +
                      (
                        4294967296 +
                        16777216 * o +
                        65536 * l +
                        256 * a +
                        (u ? d(255 * h) : 0)
                      )
                        .toString(16)
                        .slice(1, u ? void 0 : -2))
                : null);
        }
        function x(t, e, n, r, o) {
          return (t - n) * (t - n) + (e - r) * (e - r) < o * o;
        }
        function m(t, e, n, r, o, l) {
          return (
            (o *= o),
            (l *= l),
            Math.pow(t - n, 2) / o + Math.pow(e - r, 2) / l <= 1
          );
        }
        function _(t, e, n, r) {
          let o = e * r;
          return n > o && (t += o - n), t < -90 && (t = -90), t;
        }
        function w(t, e, n, r) {
          var o = t - n,
            l = e - r;
          return Math.sqrt(o * o + l * l);
        }
        function y(t, e, n, r, o) {
          return Math.abs(w(t, e, n, r)) / o;
        }
        function b(t, e, n, r, o, l) {
          return (
            (o *= o), (l *= l), Math.pow(n - t, 2) / o + Math.pow(r - e, 2) / l
          );
        }
        function S(t, e, n, r = null, o = null, a = null) {
          let f = t,
            h = t,
            w = t / 2,
            S = e.planet_radius,
            M = e.atmosphere_radius,
            k = e.clouds,
            C = e.atmosphere,
            j = e.craters;
          null == o && (o = { octaveCount: 6, amplitude: 8, persistence: 0.5 }),
            null == a &&
              (a = { octaveCount: 6, amplitude: 8, persistence: 0.2 });
          let D = i(f, h, o, r),
            P = null;
          n.add_detail && (P = u(f, h, r));
          var A = window.main_pureimage_shim.make(f, h),
            O = A.getContext("2d");
          if (C && M > S) {
            let t = M - S,
              e = n.atmosphere_opacity;
            for (let r = 0; r < t; r++) {
              let o = S + r,
                l = ((t - r) / t) * 100;
              (l -= 100 - e),
                O.beginPath(),
                O.arc(w, w, o, 0, 2 * Math.PI, !1),
                (O.lineWidth = 2),
                (O.strokeStyle = n.atmosphere_color + d(l)),
                O.stroke();
            }
          }
          for (let t = 0; t < D.length; t++) {
            let r = D[t];
            for (let o = 0; o < r.length; o++) {
              let l = r[o];
              if (!x(o, t, w, w, S)) continue;
              let a = n.land_color,
                i = 6;
              if (
                (l >= e.mountain_level &&
                  l < e.mountain_top_level &&
                  ((a = n.mountain_color), (i = 3)),
                l >= e.mountain_top_level &&
                  ((a = n.mountain_top_color), (i = 6)),
                l < e.beach_level && l > e.shore_level
                  ? ((a = n.beach_color), (i = 6))
                  : l <= e.shore_level && l > e.sea_level
                    ? ((a = n.shore_color), (i = 2))
                    : l <= e.sea_level && ((a = n.ocean_color), (i = 2)),
                n.add_detail)
              ) {
                a = g(Number.parseFloat((0.25 * l - 0.5).toFixed(2)), a);
                let e = P[o][t] * i;
                (e = _(e, w, t, n.shading_level)), (a = s(a, e));
              }
              (O.fillStyle = a), O.fillRect(o, t, 1, 1);
            }
          }
          if (e.poles) {
            let r = e.pole_level,
              o = S * r,
              l = S * r * 0.9,
              a = [f / 2, t / 2 - S],
              i = [f / 2, t / 2 + S];
            for (let r = 0; r < D.length; r++) {
              let u = D[r];
              for (let c = 0; c < u.length; c++) {
                let u = D[r][c];
                if (!x(c, r, w, w, S)) continue;
                if (
                  !m(c, r, a[0], a[1], 2 * S, o) &&
                  !m(c, r, i[0], i[1], 2 * S, o)
                )
                  continue;
                color = n.pole_color;
                let f = b(a[0], a[1], a[0], r, 2 * S, o);
                r > t / 2 && (f = b(i[0], i[1], i[0], r, 2 * S, o));
                let d = b(a[0], a[1], a[0], r, 2 * S, l);
                r > t / 2 && (d = b(i[0], i[1], i[0], r, 2 * S, l)),
                  (f *= d),
                  f > 1 && (f = 1);
                let h = O.getImageData(c, r, 1, 1).data,
                  g = v(h[0], h[2], h[1]);
                if (
                  (e.hard_pole_lines || (color = p(color, g, f)), n.add_detail)
                ) {
                  let t = 5 * P[c][r];
                  (t = _(t, w, r, n.shading_level)), (color = s(color, t));
                }
                u < e.beach_level && (color = p("#BDDEEC", g, f + 0.01)),
                  (O.fillStyle = color),
                  O.fillRect(c, r, 1, 1);
              }
            }
          }
          if (j) {
            let t = c(f, h, o, r);
            for (let o = 0; o < t.length; o++) {
              let l = t[o],
                u = l[0],
                c = l[1],
                d = Math.round(Math.random() * (f / 15)),
                h = Math.round(0.1 * d),
                m = Math.round(0.2 * d);
              d += h;
              let b = i(f, f, a, r),
                M = 0,
                k = 0;
              for (let t = u - d; t < u + d; t++) {
                for (let r = c - d; r < c + d; r++) {
                  if (!x(t, r, w, w, S)) continue;
                  if (!x(t, r, u, c, d)) continue;
                  (color = n.crater_color),
                    x(t, r, u, c, d - h) && (color = g(0.2, color)),
                    x(t, r, u, c, m) && (color = g(-0.3, color));
                  let o = y(u, c, t, r, d);
                  if (o > 0.5) {
                    let e = O.getImageData(t, r, 1, 1).data;
                    var G = v(e[0], e[2], e[1]);
                    color = p(color, G, o);
                  } else color = g(o, color);
                  if (
                    ((value = D[t][r]),
                    value <= e.shore_level && value > e.sea_level
                      ? (color = n.shore_color)
                      : value <= e.sea_level && (color = n.ocean_color),
                    n.add_detail)
                  ) {
                    let e = Number.parseFloat((b[r][t] - 1).toFixed(2));
                    color = g(e, color);
                    let o = 5 * P[r][t];
                    (o = _(o, w, t, n.shading_level)), (color = s(color, o));
                  }
                  (O.fillStyle = color), O.fillRect(r, t, 1, 1), k++;
                }
                M++;
              }
            }
          }
          if (k) {
            let t = i(f, h, a, r),
              o = e.cloud_radius,
              l = e.cloud_level;
            for (let e = 0; e < t.length; e++) {
              let r = t[e];
              for (let t = 0; t < r.length; t++) {
                let a = r[t];
                if (!x(t, e, w, w, o)) continue;
                if (a < l) continue;
                let i = n.cloud_color;
                (i = s(i, _(0, w, e, n.shading_level))),
                  (i += d(a * n.cloud_opacity)),
                  (O.fillStyle = i),
                  O.fillRect(t, e, 1, 1);
              }
            }
          }
          return A;
        }
        function M(t, e, n, r = null, o = null, a = null) {
          let c = t,
            f = t,
            h = t / 2,
            v = e.star_radius,
            p = e.radiation_radius,
            m = n.radiation_color,
            _ = n.radiation_opacity,
            w = e.radiation;
          null == o && (o = { octaveCount: 6, amplitude: 8, persistence: 0.5 }),
            null == a &&
              (a = { octaveCount: 6, amplitude: 8, persistence: 0.2 });
          let y = i(c, f, o, r),
            b = null;
          n.add_detail && (b = u(c, f, r));
          var S = window.main_pureimage_shim.make(c, f),
            M = S.getContext("2d");
          if (w && p > v) {
            let t = p - v;
            for (let e = 0; e < t; e++) {
              let n = v + e,
                r = ((t - e) / t) * 100;
              (r -= 100 - _),
                M.beginPath(),
                M.arc(h, h, n, 0, 2 * Math.PI, !1),
                (M.lineWidth = 2),
                (M.strokeStyle = m + d(r)),
                M.stroke();
            }
          }
          for (let t = 0; t < y.length; t++) {
            let e = y[t];
            for (let r = 0; r < e.length; r++) {
              let o = e[r];
              if (!x(r, t, h, h, v)) continue;
              let l = n.color;
              n.add_detail &&
                ((l = g(Number.parseFloat((o - 0.5).toFixed(2)), l)),
                (l = s(l, 20 * b[r][t]))),
                (M.fillStyle = l),
                M.fillRect(r, t, 1, 1);
            }
          }
          if (n.blind_spots) {
            let t = i(c, f, a, r);
            for (let r = 0; r < t.length; r++) {
              let o = t[r];
              for (let t = 0; t < o.length; t++) {
                let l = o[t];
                if (!x(t, r, h, h, v)) continue;
                if (l < e.blind_spot_level) continue;
                let a = n.blind_spot_color;
                (a = g(-Number.parseFloat((l - 0.5).toFixed(2)), a)),
                  (a = s(a, 10 * b[t][r])),
                  (M.fillStyle = a),
                  M.fillRect(t, r, 1, 1);
              }
            }
          }
          return S;
        }
        function k(t, e, n, r = null) {
          let o = t,
            a = t,
            i = t / 2,
            f = e.giants_radius,
            h = e.giants_atmosphere,
            m = e.eyes,
            y = e.atmosphere,
            b = null;
          n.add_detail && (b = u(o, a, r));
          var S = window.main_pureimage_shim.make(o, a),
            M = S.getContext("2d");
          if (y && h > f) {
            let t = h - f,
              e = n.atmosphere_opacity;
            for (let r = 0; r < t; r++) {
              let o = f + r,
                l = ((t - r) / t) * 100;
              (l -= 100 - e),
                M.beginPath(),
                M.arc(i, i, o, 0, 2 * Math.PI, !1),
                (M.lineWidth = 2),
                (M.strokeStyle = n.atmosphere_color + d(l)),
                M.stroke();
            }
          }
          for (let t = 0; t < a; t++)
            for (let e = 0; e < o; e++) {
              if (!x(e, t, i, i, f)) continue;
              let r = n.base_color;
              if (n.add_detail) {
                let o = 6 * b[e][t];
                (o = _(o, i, t, n.shading_level)), (r = s(r, o));
              }
              (M.fillStyle = r), M.fillRect(e, t, 1, 1);
            }
          let k = (2.5 * f) / n.colors.length;
          for (let t = 0; t < n.colors.length; t++)
            for (let e = 0; e < a; e++) {
              let r = ((k * t + e) % k) / k;
              for (let l = 0; l < o; l++) {
                if (!x(l, e, i, i, f)) continue;
                if (k * t > e) break;
                let o = n.colors[t],
                  a = n.colors[t + 1];
                if ((null == a && (a = o), (o = p(o, a, r)), n.add_detail)) {
                  let t = 6 * b[l][e];
                  (t = _(t, i, e, n.shading_level)), (o = s(o, t));
                }
                (M.fillStyle = o), M.fillRect(l, e, 1, 1);
              }
            }
          if (m) {
            let t = c(
              o,
              a,
              (cloudGeneratorOptions = {
                octaveCount: 4,
                amplitude: 1,
                persistence: 0.2,
              }),
              r,
            );
            for (let e = 0; e < t.length; e++) {
              let r = t[e],
                l = r[0],
                a = r[1],
                u = Math.round(Math.random() * (o / 15)),
                c = Math.round(0.2 * u),
                d = 0,
                h = 0;
              for (let t = l - u; t < l + u; t++) {
                for (let e = a - u; e < a + u; e++) {
                  if (!x(t, e, i, i, f)) continue;
                  if (!x(t, e, l, a, u)) continue;
                  (color = n.eye_color),
                    x(t, e, l, a, c) && (color = g(-0.3, color));
                  let r = Math.abs(w(l, a, t, e) / u);
                  if (r > 0.3) {
                    let n = M.getImageData(t, e, 1, 1).data;
                    var C = v(n[0], n[2], n[1]);
                    color = p(color, C, r);
                  } else color = g(r, color);
                  if (n.add_detail) {
                    let r = 5 * b[t][e];
                    (r = _(r, i, t, n.shading_level)), (color = s(color, r));
                  }
                  (M.fillStyle = color), M.fillRect(e, t, 1, 1), h++;
                }
                d++;
              }
            }
          }
          return S;
        }
        const C = (t, e, n, r = null, o = null, l = null) =>
            new Promise((a) => {
              a(S(t, e, n, r, o, l));
            }),
          j = (t, e, n, r) =>
            new Promise((o) => {
              o(k(t, e, n, r));
            }),
          D = (t, e, n, r = null, o = null, l = null) =>
            new Promise((a) => {
              a(M(t, e, n, r, o, l));
            });
        async function P(t, e) {
          await l.encodePNGToStream(t, a.createWriteStream(e));
        }
        async function A(t) {
          return await t.data;
        }
        function O(t, e) {
          let n = t.width,
            r = t.height;
          var o = window.main_pureimage_shim.make(n, r),
            a = o.getContext("2d"),
            i = (e * Math.PI) / 180;
          return (
            a.clearRect(0, 0, n, r),
            a.translate(n / 2, r / 2),
            a.rotate(i),
            a.drawImage(t, -n / 2, -r / 2, n, r),
            a.rotate(-i),
            a.translate(-n, -r),
            o
          );
        }
        1 == 1,
          window.SteffTek_planet_js ||
            ((window.SteffTek_planet_js = {}),
            (window.SteffTek_planet_js.asyncGenerateGasGiant = j),
            (window.SteffTek_planet_js.asyncGeneratePlanet = C),
            (window.SteffTek_planet_js.asyncGenerateStar = D),
            (window.SteffTek_planet_js.generatePlanet = S),
            (window.SteffTek_planet_js.generateStar = M),
            (window.SteffTek_planet_js.getBuffer = A),
            (window.SteffTek_planet_js.rotate = O),
            (window.SteffTek_planet_js.save = P));
      },
      234: function () {},
      603: function () {},
      462: function (t, e, n) {
        "use strict";
        n.r(e),
          console.log(
            "we are in the browser. No need to do anything. Just use new Canvas()",
          ),
          (exports.make = function (t, e) {
            let n = document.createElement("canvas");
            return (n.width = t), (n.height = e), n;
          });
      },
    },
    e = {};
  function n(r) {
    var o = e[r];
    if (void 0 !== o) return o.exports;
    var l = (e[r] = { id: r, loaded: !1, exports: {} });
    return t[r].call(l.exports, l, l.exports, n), (l.loaded = !0), l.exports;
  }
  (n.amdD = function () {
    throw new Error("define cannot be used indirect");
  }),
    (n.amdO = {}),
    (n.d = function (t, e) {
      for (var r in e)
        n.o(e, r) &&
          !n.o(t, r) &&
          Object.defineProperty(t, r, { enumerable: !0, get: e[r] });
    }),
    (n.hmd = function (t) {
      return (
        (t = Object.create(t)).children || (t.children = []),
        Object.defineProperty(t, "exports", {
          enumerable: !0,
          set: function () {
            throw new Error(
              "ES Modules may not assign module.exports or exports.*, Use ESM export syntax, instead: " +
                t.id,
            );
          },
        }),
        t
      );
    }),
    (n.o = function (t, e) {
      return Object.prototype.hasOwnProperty.call(t, e);
    }),
    (n.r = function (t) {
      "undefined" != typeof Symbol &&
        Symbol.toStringTag &&
        Object.defineProperty(t, Symbol.toStringTag, { value: "Module" }),
        Object.defineProperty(t, "__esModule", { value: !0 });
    }),
    (n.nmd = function (t) {
      return (t.paths = []), t.children || (t.children = []), t;
    }),
    n(44);
})();
