define(["require", "exports"], function (require, exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    exports.config = void 0;
    // TODO: ゲームのパラメータをリファクタリングする
    exports.config = {
        maxLife: 5,
        maxSize: 100,
        bombLife: 30,
        bombSize: 30,
        bombSpeed: 20,
        bombFire: 24,
        bombDmg: 1,
        missileLife: 40,
        missileSize: 50,
        missileSpeed: 32,
        missileFire: 36,
        missileDmg: 4,
        dmgSize: 12,
        missileCharging: 300,
        missileCharged: 700,
        dmgMessage: "Ouch...",
        missileMessage: "Help me!!",
    };
});
