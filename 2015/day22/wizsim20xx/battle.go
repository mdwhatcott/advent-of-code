package wizsim20xx

import (
	"log"

	"github.com/smartystreets/logging"
)

type Result struct {
	PlayerHits int
	PlayerMana int
	BossHits   int
	Turns      int
}

type Battle struct {
	log *logging.Logger

	Result

	PlayerArmor int
	BossDamage  int
	WizardSpell conjurer

	shieldEffectTimer   int
	poisonEffectTimer   int
	rechargeEffectTimer int
}

func NewBattle() *Battle {
	battle := new(Battle)
	battle.Reset()
	return battle
}

func (this *Battle) Reset() {
	this.Result = Result{}
}

func (this *Battle) gameover() bool {
	return this.PlayerHits <= 0 || this.BossHits <= 0
}

func (this *Battle) Simulate() Result {
	for isPlayerTurn := true; !this.gameover(); isPlayerTurn = !isPlayerTurn {
		this.Turns++
		//this.log.Printf("%+v", this.Result)
		this.observeEffects()
		this.doTurn(isPlayerTurn)
	}
	return this.Result
}

func (this *Battle) observeEffects() {
	if this.PlayerArmor = 0; this.shieldEffectTimer > 0 {
		this.shieldEffectTimer--
		this.PlayerArmor = ShieldArmor
	}
	if this.poisonEffectTimer > 0 {
		this.poisonEffectTimer--
		this.BossHits -= PoisonDamage
	}
	if this.rechargeEffectTimer > 0 {
		this.rechargeEffectTimer--
		this.PlayerMana += RechargeManaBonus
	}
}

func (this *Battle) doTurn(isPlayerTurn bool) Result {
	if this.gameover() {
		return this.Result
	}
	if isPlayerTurn {
		this.playerTurn()
	} else {
		this.bossTurn()
	}
	return this.Result
}

func (this *Battle) playerTurn() {
	spell := this.WizardSpell()
	this.PlayerMana -= spellCost[spell]

	if this.PlayerMana < 0 {
		this.PlayerHits = -911
		return
	}

	switch spell {

	case MagicMissileSpell:
		this.BossHits -= MagicMissileDamage

	case DrainSpell:
		this.BossHits -= DrainDamage
		this.PlayerHits += DrainHealing

	case ShieldSpell:
		this.shieldEffectTimer = ShieldEffectDurationInTurns

	case PoisonSpell:
		this.poisonEffectTimer = PoisonEffectDurationInTurns

	case RechargeSpell:
		this.rechargeEffectTimer = RechargeEffectDurationInTurns

	default:
		log.Panicln("Unknown spell:", spell)

	}
}

func (this *Battle) bossTurn() {
	damage := this.BossDamage - this.PlayerArmor
	if damage < 1 {
		damage = 1
	}
	this.PlayerHits -= damage
}
