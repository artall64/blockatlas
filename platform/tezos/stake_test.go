package tezos

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/tezos/bakingbad"
	"net/http"
	"testing"
)

const accountSrc = `
{
  "delegate": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
  "balance": "91237897"
}`

const validatorSrc = `
[
	{
   "address":"tz2TSvNTh2epDMhZHrw73nV9piBX7kLZ9K9m",
   "name":"Polychain Labs 2",
   "logo":"https://services.tzkt.io/v1/logos/polychainlabs.png",
   "balance":2572632.898746,
   "stakingBalance":25743192.663017,
   "stakingCapacity":27842122.223558,
   "maxStakingBalance":27842122.223558,
   "freeSpace":2098929.560541,
   "fee":0.1,
   "minDelegation":0,
   "payoutDelay":6,
   "payoutPeriod":1,
   "openForDelegation":true,
   "estimatedRoi":0.055564,
   "serviceType":"multiasset",
   "serviceHealth":"active",
   "payoutTiming":"no_data",
   "payoutAccuracy":"no_data",
   "insuranceCoverage":0
}
]
`

var validator = blockatlas.Validator{
	Status: true,
	ID:     "tz2TSvNTh2epDMhZHrw73nV9piBX7kLZ9K9m",
	Details: blockatlas.StakingDetails{
		Reward: blockatlas.StakingReward{Annual: 5.5564},
		StakingBasicDetails: blockatlas.StakingBasicDetails{
			MinimumAmount: blockatlas.Amount("0"),
			Type:          blockatlas.DelegationTypeDelegate,
		},
	},
}

var stakeValidator = blockatlas.StakeValidator{
	ID:     "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "stake.fish",
		Description: "Leading validator for Proof of Stake blockchains. Stake your cryptocurrencies with us. We know validating.",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/tezos/validators/assets/tz2fcnbrerxtattnx6iimr1uj5jsdxvdhm93/logo.png",
		Website:     "https://stake.fish/",
	},
	Details: getDetails(
		bakingbad.Baker{},
	),
}

var validatorMap = blockatlas.ValidatorMap{
	"tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93": stakeValidator,
}

var delegationsBalance = "91237897"

var delegation = blockatlas.DelegationsPage{
	{
		Delegator: stakeValidator,
		Value:     delegationsBalance,
		Status:    blockatlas.DelegationStatusActive,
	},
}

func TestNormalizeValidator(t *testing.T) {
	var v []bakingbad.Baker
	err := json.Unmarshal([]byte(validatorSrc), &v)
	assert.Nil(t, err)
	result := normalizeValidator(v[0])
	assert.Equal(t, validator, result)
}

func TestNormalizeDelegations(t *testing.T) {
	var account Account
	err := json.Unmarshal([]byte(accountSrc), &account)
	assert.NoError(t, err)
	assert.NotNil(t, account)
	result, err := NormalizeDelegation(account, validatorMap)
	assert.NoError(t, err)
	assert.Equal(t, delegation, result)
}

var mockedTezosResponse = `{"balance":"19924380870","frozen_balance":"16854253212","frozen_balance_by_cycle":[{"cycle":240,"deposit":"3072000000","fees":"15676","rewards":"90000000"},{"cycle":241,"deposit":"3328000000","fees":"621522","rewards":"95000000"},{"cycle":242,"deposit":"3648000000","fees":"4360","rewards":"101250000"},{"cycle":243,"deposit":"3392000000","fees":"28321","rewards":"125833333"},{"cycle":244,"deposit":"2816000000","fees":"0","rewards":"55000000"},{"cycle":245,"deposit":"128000000","fees":"0","rewards":"2500000"}],"staking_balance":"163002104470","delegated_contracts":["tz1YGsmfUfed8fmG4QLhGB1Y3gCMX44t1JTc","tz1WnfCXDniXhHzbXEPTQ1JSoyb7EzotYJs3","tz1RqtpgLw3uedpWaDVC5WAdvNfiUnsY7ZC5","tz1eLjmx4j7QLX2RabyS8Pf3sSe1U6H2UZji","tz1TYhMcdYah6QJWQbL853x4Rn6fSvtHhqDf","tz1XcmqCntUEzrpV6xcvw449Ro83KU5W3vv6","tz1QfbGBmr5dvha6agPcmsvnaz2EfcctW4ZJ","tz1bN4kotiEE7dAEsTGWpNTF2DnjJw9NNSoD","tz1hoj3Lw71qgsQf5a7cdhdU9crEH3mAAaNn","tz1SfvgWvSBxNsgta6UCmytMTPHzMsmjMB54","tz1ZsKynrRs8KNe3ao1ZMaKk6he6VUWtjcFU","tz1ZEKCiu7QspFVZL1svyr9vvir7tTm5CEdW","tz1gScW33c2EjpusjaPCm7r82QVZsALnNsXk","tz1TDnm46TxwLRbihJf2yTQLJGJYe6k7N3Kw","tz1NQCizMgAvmjhMJEidnTNPFEPHxA6XLY65","tz1dp6YqFkjoDZqrceSEC7cKFDEWbynyv7UV","tz1RctGdDVhQFoqL99bXJvYTwQNJWEL2Rx6Y","tz1hqd37vwxZcNeiNMrz6kpujusPRxcaFyhv","tz1crpuStZ66tv89kZMQbiGHk2C2RyD87o2S","tz1a45AtW3gs964AxC7X2sShAj9wsBGQjsWQ","tz1Xi1AtnoAv74PAQFDUY67TEjww1J2e6FaQ","tz1XqSWe3M9eEd2mueqEWmgKZLoYVJCRC92B","tz1Y5JZMUEsuX9oicG9AqXG8XWPHmKeWFrb8","tz1LVtuGkbCgEZsZhbzu1sonVg5xZyrrWqGD","tz1cfqx7H6qsXgsLTHFn7hWjdkU1SjpzpFyx","tz1SFMBqUwq6JJP1mQZfp8ae4GsfGSM8812Y","tz1LeU73Tm2eBcUV1prcHLjK1Pb7KgC8rPTu","tz1dMvyK57FBfUqzSenaWSoY1Rng77KxwTT8","tz1S382giorXLVdzw79LVcMj5MLJkYK8Ex8d","tz1TAscJC43xhPBNBESuNupw3FhLsfE4Ghxd","tz1eACKAsqr5faLgsE3ACdHeejrdM4igY2qj","tz1Ln7pr3Bf7Yxmnf57iEkgEUvo7PrMh3HJh","tz1XC1DbtQjUu2udaeVbSa3G2rrU9dJKmN6V","tz1PCa2Drre9MaJ5pcKpQHh8A75LDEYYA4rq","tz1S7SacT3Pe8EfDfMEnR7fdvDamZDiwVMoQ","tz1U6V2mWc6zd31yFXME5Xnoifo3PeHQhuJA","tz1V54Uhpp2i8pUVEFNcmhrkJqfDZhqb5F5C","tz1aTmVHkGBChKHfAqdJjjadkP4rz4MDjbJS","tz1dpijoKv5dCFT1BvtEmxy4YLcCsAc9ZGeh","tz1Mmv9StLbPu3jYVBDoKgNJTb3FD5wuA2vA","tz1PkJZs2Eb3ouLd9cTf4wSQ3n9aMz7YLGNR","tz1YpTsHeubr7WQaKu3Yfs3m1gKw5QP1MbdM","tz1ZP1EgEkNwB1nYA8bQdJ7YcttsbJuJCZLe","tz1Wh7XciNS6oUTW14bQ5PD7pQjqae2tNA8S","tz1bgRKKwMc4s8fxTDJrauCx3NoGLSRJ5B49","tz1URZ8CPFhzLz3UjQD9zUjTZY38XwGTCiKn","tz1KrUvC8nUGRnk4RXtqAFLsJj4cbtXjR5Wg","tz1f8eDjy48LDcWrrNTF3wwkGEjyXqDcoMro","tz1RStdCWJXpVVvYvAB5VZXgzrxPGoHpEhjK","tz1LWjsR3LiBDPyvuf7S1qJxaVKNZV9TSbaW","KT1DMboG5KtaPbMTaA2PeV8FqRcqbt9adzQE","tz1iq1eA7i2ogVCQ52g3zeSEqCrFTSVxhE2c","tz1Y9tfG6Gs5oFdLtcE9obQaxbwZyK8CTn6e","tz1XotDGePt679YtqexwfFvE9AutjdkvemH2","tz1LG4KEFVjp1MxSCXQWUSyoZpE4Nq3gbxHy","tz1TMRCpNgKAdp8GH3jXAVR3MuLMMPZBzevi","tz1hZY5ix3GTTUMAsv6anwD2817RspnnEb81","tz1gzPsf717fNkpFVQYuNAKBBQngqGx6CFJz","tz1TTxhXmRRdMAKdF7ykdyybgg1baD686JWs","tz1dPnbKhnmJo3J7S2rokMSmTtaWNGDeQuYa","tz1irRBG7qLAeo6URDn4bnUb9NByoLQUHN7K","tz1d9d9brjpFRrErRF8tT5StDve7svAQueTB","tz1SBRVFKJCaDnPRzHrn1Pc7NG61tqNY1HpM","tz1h4Ag9NixUomVaFo8MatDpEL84JEEs9oFx","tz1NVzxdVF2gy276bGAt3TFJHWUMtnS5BZrN","tz1Ytb5C7xDvWArsif3cEfWC7NpEiWkGVcW4","tz1WQ4w3WQUzReKAehyDKkVQjiVNJLuWoCrx","tz1StPqh4Jybs5yE4adLLBvrirwm7cuoEMC3","tz1f5yKJQwdEx1XbqUreT5E5MG2ddoP9sxG6","tz1hVxLkK7fDPCsgRytKyqSYoRGdPYkRAHK4","tz1Qs7gcTNMN759qBUukQmfA6cUeWzqtU7sq","tz1SsR67T1ykErBtVinVMCWhvwrUWBMkzxQu","tz1QD8a5kJBGHArP1DHECBzj4gbRPZuBGGL9","tz1inQ6S764Btu5HgHRGMg72Tee5UjQtri7y","tz1asvE47tkUoYPrb6HrVzkHVhfx3roZA56f","tz1ZqdsHNTeYLCw3Etw1BwMTRbU3Ji42HT4N","tz1Zd43K6xBnpddV2pv2PcbKPBXpYo6qAAxc","tz1YbEBGE1KQXAMLiQEnXVKLoNKVAmtQxJ2y","tz1QFhih66LJDFpm9T6a2gWdg7551k65C9ke","tz1T4RpJjC671WadSwA145zMZNSCmvcwGL5R","tz1Pibvp7EttcPX56CTvmxPsSLqwsDhd2fC1","tz1cFu4B4afDVkcojpv43h7wbv8169VA7mSU","tz1ixKgGdRoSzsp57qe5WXHLbeu2GRpXYmGW","KT1M6T4ft8UXXwFxa6QxRU3T1JDEKQBRK8Li","tz1csRh2tzRZw8dcuAm1LMCMUQBEq7N3EmY7","tz1gaz2QgwUTyZypSmMx72oUrraD1hZ4CWEB","tz1TpW2WB8Y9mYsDp2Lx4EoKcocyFWApScFy","tz1f9Hg7KF6xKLabL4UNLPdtJeohQbuJrVa2","tz1Q8Q7sT3r7pdGGyYEtX7v6jTxCwJ3bSdtJ","tz1iLAxh9M6xcaNMVhb5hGxZb9qvzBVwo5UT","tz1XXHuK4tMKJZBQDeDFCpzWdH4g26Q42gFW","tz1VvZYohqJvKzTKsRYcfuQuak6LYFao9hkD","KT1E6sM9KBpYWoK7ZxoxrcAzbngVhn9ZxkPw","tz1TuBtRJ5bGa7KRyhiZcoFiZTHXnuq4wz4X","tz1fQRtYbQqenvb171siK8KqScBZKHrsqKW4","tz1dn9tn4UqhwLiAShcaq63oiwzQ7X8aySTz","tz1Uvknh1VNhYYQbhFaP1rDU1GNtFpGWG98c","tz1ZosT8MD6tibQDYLbhfTm5Z9oQPL2f7LwP","tz1Yk9Mz6hSBrq7qWLUZKDqcuwyu19i7FS4Q","tz1i5wEnEp1Wf4WiL7dFxiTj81VSiDprGWyZ","tz1fQv9BGCjZC29DFPSmqNxkjFx2tLzRy9Kn","tz1M7TVpisBJfPkkjrfsnK7CZyRbVb2Vdooa","tz1TJ6ZPJVrAFZUNhReL9FnhzXXd86Hmtk3d","tz1hbeGj3krpLMdKMumeukTatZcDCG9RGnyT","tz1Xm41xUyqvV5caWkpx1ym2p8uTvJPKZMdw","tz1XZEpd7EnfzmjYKfRY2zrfEd4trBJWBybN","tz1QZAQoaXbLXW3WjYB8os5xjwe4mB62KjPi","tz1gd8XNS2Me4oRXq9XTkTX9kYumRM4Y6rXf","tz1aK1rKDWDSa31kB1SwP1tD3BvPTNQWXwwK","tz1Ui8BtvQpJpF155ppgMZSEpCKEDuLXNMX4","tz1ioGHFTom3ojBbcfgEb8mt1f1ceobVEcEq","tz1LxreKnQEFW2QB3dZDX1bNThcW8vmem4E8","tz1gu2wRjBfY7jpMokMDZGNZVBE5n98J65g4","tz1dfAqA9rhx1tFuu86FdhBwKMwaisVe1dEg","tz1ND2YvsV6mnWSUHA2bd1oUcTmshMbxYV1M","tz1aSBidSCjNqLcEk8BYvA43EenEZNtsvkuW","tz1e8Fk5whgxrPzTZAMiW2FbatdaABNt5hpX","tz1QeU3SH1pGYTWmKfZKovH2RiiqSnd3zFPo","tz1UkUgyb3Au9GeQx6YbcJJaafNDpgGRLDeV","tz1WPnZxticG6xGpHNhaV9kvU1PfgYiMXp7A","tz1LKVTXzQtDwQWXJpbe88Q2U5TceBrEJf5i","tz1X4tzFR73YRBBora46aevdGn3aqz5UzdrJ","tz1PouWF6hyoHmMYAGuwDBAXJKL4Z8yixEwq","tz1ib38e3VCZgvRxc2Gj6KtAgAv9FkWx3xt1","tz1Ndenm8uaonPs8oo7iqXQ6RRho6cG8GBd1","tz1WRAfQ4c98Lxz5k2j5QFwcaw31Z6zSuiza","tz1W7G8pgorMWDcWSsTpQLuQoKqVPk6MNTV5","tz1NY7t1otGmLV9VuhdApzsqwwqGEK6iLeyH","tz1MA5ywEYaPNzJWeLWBDfvAfZmryeFdrexy","tz1Rzq5gz824wG4gt4MPX13oKojnY3cJVTAB","KT1MmHkxDKy16ETe6CKY6Hs3SXGzAeMKHcmy","tz1aYL7YjKn8zFV8kCNMg5qjJTXQ5sPhnVQU","tz1Mqazkg2SgqB4ibXxpFudp2HVYEZ2s3t59","tz1X1o4zSNrnTDouETFM9AXQzamWcUkEj59n","tz1Qe9fqiQnizh5gzuURT3J24MQxd7A9qxTt","tz1UWyf2N1yCiEkfyCKdyMitKMAaUWACWpro","tz1LYCQsxVf2RdDEWcG5N1zY72XskyXevciC","tz1XD2yz8UQYGJ6BRb7h7wZp854fQXC9DQFm","tz1iiT4XbLvffTJbGhuiCg3ofaaNdLX2thSu","tz1YrCk1KZu9wwn9vCUYuZp8medpViizrZwK","tz1QPmsd6BS9CUAugcNGJP9me8V6qtEEYrV8","tz1adgGe3ghX6DH7kh1BAasaCGrf6ermGY53","tz1M7PNxBBCUx14jfdv5ZERWZWBhPJVjKHJD","tz1VcWpMRqd5xxoauw1urcnPQ2kZZxhbqgcv","tz1Y2nVkzQQmdVnDcuo6VDAgnBufoUHLXR3r","tz1bGvZzpjJZHzFJGQYWPVg4JQ7arsmX3CBW","tz1ZGn1qwM2Bg65xqNx37XB4ptSQnRh7LjED","tz1Pku546auwaUBm4Wht1zwZtfxg7ic81p6Z","tz1R5yUME8wSYtyZNpfWdoPZcQgNLtaUUvef","tz1b5HuS35FKB7qWvFuALXX6ZFspEWWJVjxP","tz1ihSZ12F2y6H778E6L4ciWzvX7rfPwmGQd","tz1T4yyuchdPqtZSbtsaBS3GLqLeUg8dJmb4","tz1LnRcJUtFBKcht4n2jo2PAJF1ufwAVa1nL","tz1iwdFtYWaicwc1dfJA6YmN7eRCX1RkaGuz","tz1aD3HxUFLSr6tmYuP7HKvx7t8WgKTY5pYs","tz1XJMJ7XjbHkkFkAWw9C8ocjfmkmTBXo5TG","tz1PfUUZ1YTRqP7qkjsSR1X7TCeYbRpnCfaj","tz1aLVDxteKdmBRaRMcQPQaEL2ExbDwQA7fP","tz1QoFWg86UsKYjtskkmBkz58mftgHLrvrci","tz1ZWbmnG1sb4KCkrjii1FoiJ1RcJEg7Tf5V","tz1VGrdEoiWdSpUFGZAbX6hxuQAJjEXqXvwW","tz1Q8gvcfGC74oaWx2Gb4TipRLb7aEfsWyAG","tz1QVsVf5VQfihDiBED7ZMn5Ai637z9GwfpB","tz1Nqf8Do6NFXncMzGhYasnif7iTJ8UgGHTn","tz1KotnQg3tbzTLxYJc9XW3dagmYjAMHzYV7","tz1QzMbAcThy4EX8ENvHZm2BVxQMmCY7wsUw","tz1YKTfdWtb6gzr26mRsLsGCNK8TBBVEHC1E","tz1egRyJLzTmC4aMUCRZLKEhtYd2MyD9f8RP","tz1NJFR2PLvbijMGGZ9LnW88CZxo1SSFxsxR","tz1fU6Zxx14MCEqfkSwDWRMqFmYmhygSJrcK","tz1aN9Gp5FXc7CQoR8MGSf3VUFohe6f9rFJ6","tz1KuQ5yDerio8tW294zHsYofWc8TQ3xnbaJ","tz1Yi1qojMoS2JhESiE4vAeo7dWUqQbnEQeL","tz1RxecATFfopVUvE9HS45rL6rgzGVg4cFG2","tz1hXRjun9ou7YN7sFKJJsgHPkGpdXzamzKY","tz1dWg4SYTFiYfPrDGBNCaNZunjCi9CyK8PU","tz1N7CRoQTrTjM6BG6PK4XfwnkEERKj1GgTt","tz1T8RkwAcBBwc9uCovercc6Qr326ezqtoij","tz1UkDYB4RgqBcRL7yYSikc8r6fpAeRVP9c4","tz1WLjJEwxQK3hbnZWkmEzGzQ988MdQVq1iB","tz1bLijXE1D8XxHjzK9wBK2udxNqXbHkMzFg","tz1VMWRzCPN9E7NLBfSHxFVzjZ2vxNPKRXCT","tz1eUfNmzeZWccsAV2ZzCovjMR4HSDH6eiiq","tz1VASv4WieYkgmUCeh2j5ftEw72NkkuCDje","tz1fZH9g6npSMZgSGzcaqegnuftJ1AcNEQ6F","tz1QGzvcCDezxV439D5cjRaTFXC8Y2bF6qTw","tz1PxPHvFdkGY7XkyuV2eUDYH3wxuerKiYiC","tz1YJUngcjMCStvP7MABMVBPyFkU3BupNymQ","tz1NoKTHMvzPooEbyeMHQBaia7Ekf8aTrv8g","tz1aQHYPPW8dq7ucWUccV74cEyzt6AiMjrfj","tz1dpCdky48ELUbjrzY89Uy4fTukmmJdr9tJ","tz1SQcKWYJCatD4jVFRNtR8UVBmtZT5VAM3x","tz1a6rVyLh6z7akANZsBuxSb2g9gNthXuFM1","tz1eNv53m15QQhKYyREahAx6LCK9YYczEUPe","tz1Z5Lgu8CDhbWNdyEzXkCC2TrkHrSDVwMWp","tz1Prq1iQE4xMn3JqwK6rMaJXoGjS9sfWM3L","tz1hxtCEcD7idQJEDiJEq37vBkoRcwF6KC2X","tz1ZTmAakWGa9fefwoaomG7wFPQhPXepyxdT","tz1dXKPAsBG4DG6oKki84UVTBYw5doTGbu95","tz1ZGCukFkDD5ns2aA8HuKGrE68DVmbhLvvs","tz1cvSwVukfoTNS16f1Y12mHYqsrnCK2z7ZM","tz1PV34FXaBY5mqLLyjG8TdzNCo5SUz51d1a","tz1KxhXker6qezyPyzpLcTVXxAadWuaRMB8Y","tz1cZmvFJgdJfUoBF2USegkoEq653qx3KgxB","tz1M8NxQupHMQ5uW1k25wCYYdxMWojKfiVNa","tz1eqnkA7YZS5uVhH42WfZkwkECTxCkUqKJG","tz1Z2ERSqjQjTtGdv6XZJ9ZoffxAQVCLHrLM","tz1iXW4XFVvg1QS6bGhuPo1VCGT6WTgEaoA5","tz1VMDp7io6vgXo8jsLcgbsUB6Vc9mdAw69x","tz1edVs2PENzexqcwWv22pwu4HfcgXbKcV4R","tz1LSUEVASTpic9FawaAMtZgR39LSYmPYcLM","tz1REnhE8TEzXdcWSRrzBYPWuveyY61Pe79R","tz1TrDkd3wtJcYnXZiB1kdoHtDJ61eP4QMA2","tz1fKL8yNKW98UGCE5duxSyM11QqzFBvQW36","tz1dPXJpxXqqLeoCs3X4GfribJKfyFVH1Jta","tz1NhzBdBcokEFPuzfdMrjihST3WU1F7uTGe","tz1crd1XnXNsCfgXVaHzYpsAe6XBEuhcYRF5","tz1N7T11AEwWAoqE6pYyU8ZYcdfBnMyPmr4X","tz1f5V7ijjQ8WBN5UUAnQvGcdXg7cAPxGCUu","KT1Vhqza4MMuELxENaiBYc6UB7syRpaZ12LF","tz1g81GLa2GeT6AgQ3CxtYBdwD3kwMLR1Hpt","tz1ifm8828dvFuS8wbwqCGhyoGKoXMAkpqbz","tz1ZfX2axfXid7Bnt2BiSv6QonweVFuuaAhM","tz1aJtqVgev1Jrc4utBfgZDj83ZFLjJKiZJu","tz1TyyjZ9adDW6rranHNXQVqEMp4Zb2rhegW","tz1iZS4BRLqyLbWbyyKt3e9gvAa6MJnDM2qj","tz1XM3wst4ioyMBHdkCVMdo3nfqYzb6zqbpn","tz1g54U8WwX4h7sVAfwNAMfmN191E1nYD42V","tz1ifvWnxwudcmQptnvbb5aRyDyWLLdUJmv3","tz1RDNwx22yxJ5x24ZxsrqFtA4bXY4gtwm4Q","tz1YvE23HCvxMNPqWsgQyYMw8Rm2UeMnPCMs","tz1dYeNiX8ayqUAc1SzDnr1iy259L5Djf1Rg","tz1Xrk1gZ6UGs2zr1RBCgexckHpS8CjgtdX2","tz1cp2AxYptGFABzTbnEWo62Rhixc5HKpjFu","tz1SCWJXrAiGQmoBPMLGifjrbTi7HnKgEDdu","tz1fxbW6z9xts7LnbXqvcdoqwp13rBavG6hP","tz1imF1LwtUxRojZSK6povb8mePfxQyk7UiF","tz1faBfWwvmovgkngNv2v4bW313jrepqdjAW","tz1XPjXcjWvqiNC4erXakoCiMY9Xx3omBRvZ","tz1iuPqyGQepHBWiVgTTdndbpnfkqckQb94n","tz1Kuu9FQu6L4bkNwLt39BGWAMSfDefWMCR7","tz1eEea8fmMweaMGbtUSa3AVWE8pgC7WM5Ym","tz1ceedKUzespKvmnaZApaXXcZTDJdVjenBa","tz1QbVugeaLaXJXqgD9s1gAYMwAJkW9rTgn3","tz1NqVdJps7ppiy5EbahVuAmDw1iKjhHsXG9","tz1UJ6u1bGV182edaXg5uPdbJq1qHR6xdJSZ","tz1fQiLH4SyuXumvKejsnQZgdnJ33J3kCqk6","tz1SBTErzCAU6o85X4xTdRzzCHQ79JkmGgRx","KT1UXH1LJ33HSaEvzW1Pk92tLbGEdsqcN93q","tz1iRk79QEsJ6Lj5F7pGaTAiQKTYw9RSUqzU","tz1ZKvoYkCB8otTqdWp1Mm3WvemXMQCSpsgT","tz1R6esD6XuzJ7hrC3ngrWdafh9FFJ9NV3oT","tz1bsVgtu9zggDAMvHWPWZET1wMEMTfzDTL1","tz1ivnzrqG3bpXuDchcS5KCpQMvEvB14Wxw2","tz1MaKRHHCaUqfaY8SgiHBzNNgvUF8EiB9RE","tz1LN4xvgHB9dqiuc9fqco1ZEwDueXDgDw8T","tz1i55Li1zdF52PXSKGavRjmE6pCpSgmA5To","tz1V9wpVevfkEfNe2JnsBGuzQKJfpGzvs6vB","tz1dGRzqFXkur4d4HEr5q5ULSTpxtT8a8y3t","tz1QVCn1BYxpeWSVsu6ZvoSm1vMwFs3U49ZV","tz1Ltc6HqhjcMEXGnW8wSaqEJfi4kse77oHS","tz1LzvmQDVt7iTDevkAu6TZAd6LkhYaw5i8W","tz1cGaQgUitUa9TYfDdgAKLY5zyh5GUVHPzs","tz1YMFNXmBU81a3hkDAf5w6oADRPuieefX4i","tz1QMMwLkiHSRxXMP2LUuLfvw4HCq2c5JfMk","tz1UYdeecJh8w5EoraCHT1TgKA1bL4jGaDWJ","tz1Lor69x8DPG9AJ5xSthwVqQ8XtqLvnSofa","tz1NUHJF4PcYq1QbN3sW7RX76b81r96hKKoz","tz1YEhhVsJTQkxcgD24tQEjfhBKETkdBip9H","tz1SVDQFfuBiQV8qbhjw5YGxxVWFkPvEmodP","tz1hjZuttQSUcSbkg47i1izUMykXCcdqq1zx","tz1fF85RhHDzszowsTj9KGwcFwfDjfWqX7UM","tz1LsXSAtqk2UaAi5UqPgnxiMSx5aNarTiZt","tz1ZMLcTZuhe1FrCtbbTyiEp4r7PBbQZ52HQ","tz1dLBPNqqKTEqWEMqyzh64jdTVMoaZij9sm","tz1VvsF4iSFq7KVV4YFznaN7ZQYSNqcrQbyS","tz1ZbdR5gF7VrGT81o2Ny4c9UVfCc7gzB8En","tz1gC8bBi8dJKzEn97fvc1cp31qBctXUHfPv","tz1WsYZt814Mg9brcP2scWD9CcCywj8NnRBQ","tz1fv4smnQUbKXARxBKkUh5tQtADqFMKKQDJ","tz1i3UYmtRdoDwcHFSRV2e3Tk47fchbwUMwT","tz1XGjuyXeeCK3EuHYnSVMqqeJj41aNDzGDG","tz1SgfF5YhWA7Y1L5kEwv5cgRTbLm998fJC5","tz1SoikxCAviVXxs3pdc8UArM6unU1mffMHV","tz1NnvH3F8SfaB39KvD7Wyep1ytDCXjPoCJ4","tz1VCjP2P3sBtCLQXKkDG6xRyz1o6Sh4cngZ","tz1YqLMuu5gqPcWZjGAaCY4cHgKFW1bUYeKR","tz1Nev8pcv8qZDK2q2Ps4isYMtcJpXWaGcUa","tz1RrEVWkWxxQY5uLuBbVU2WR9W3EFtU8iKf","tz1dfc5ei2TSihvEpFnyTTb9nuQbDWvEEEvN","tz1iGL8vPK2DNvYf67FrqdsTDyXCtpBMbgkC","tz1YXFAkEgAbgy5Mrsr2DnPhbuq2WnBhnDCa","tz1SJLqfbEyGJ21wfhMJvxgu2EjQ346o22mZ","tz1NwPTbuzs8s6mrGEJisP9CaXC458jjzzPF","tz1VGdetQxEuDWGQYE5BDREBvpzQY9javJyp","tz1awa88jNvcsmbBHXbGLk6fcqEycYZecic8","tz1dJKVHdNapMkayYcmHFVYutVXXHx5mTKri","tz1NneEsPdJRrDgse4TzLpRqvbksphYBXBMY","tz1gtGWmzyGZTz7D1BheVa4ooocxjwamba4R","tz1VXvbmwHskMRHcbcCLi6bihPt5KNtMgVeX","tz1T5MBz5WXy4mCrdVHy19zrFDa3uu5vuKG3","tz1TH3yto3yViq21nG84zmcTmuvRfVgfMRGf","tz1eJ7HaiQu85aNYwXnYByr4voQJe65wWTzv","tz1Vz1XhyYidM2zjSZoN2H5Ku3a4ctHVvGm7","tz1MVTqQAvvurF7Abs4vBmPUoSFtqPuXacmv","tz1hFFvzPH5DxuHA4ivQvrtEX92FGntGMRBi","tz1fmLPPwGznV4R6vFyW8sswUZYwcYLJeYAp","tz1N6Xh2h2K4unzgXv1XitABtA3eB26wDAUD","tz1iGEPSzjQN8PzT5FrGUzhRQnx59hVzeXn8","KT1Prk9xaXSd7uz4NzWQfZ8bWyY4csVvox5w","tz1et74jh4CLmpfmvYCG5398iygUkdL962fy","KT1PHojJp48asdfyvT9kMmnsEzB6cdUGpwFf","tz1fNt7sqTh21sekaMdsPTQvjYsC1SduMagV","tz1idn8cD2bFaftnFzvJ2gik9wGuWKDGmhsP","tz1Qz93toMBURzKyjbECFe8YKAwARu8t1C3C","KT1FBMtK4T7K5UC3nuc8YRG3wnQK3oaCH9Uo","tz1Ly6UMQm3NCmDQ2udu6WvPCJGqauhR9FWk","tz1epXpov9hcoRUaGVJ222t67udd3vV6jn35","tz1ZVR7QMhJWCKEhp4SdynXNnikdxTs5XxQT","KT1KRdu1hDLtDRufZqNMU2z42uKrXhmsmzJd","tz1ivo3rkUgJQ1Nbe4VEhNetCREAvbGZwi25","tz1UBQvZkovRpZwML86qTpL59gedS1FeWpRw","tz1d9CzUNaFMmsnUZTihrz6wDA8ZP3GvAKbj","tz1MZHrRbLDMczK4xtJP1BrRdoYnoMq6c3P7","tz1Rw3VKh9yAWLWaQMhzWoRRQoEouGTCzgV7","tz1iABqHT2J4XKEnHp8JYJNiXEJtFg8r9RbC","tz1fxdeeUWT6YG5aSQcaJVU5sbeFCQYKCvPX","tz1dNA8PnEfRgV4ZqrX3i43Lh7Rxpx9bwy78","tz1MEEdk9w54LTcHz1oSkwAY9UUSepg8WrFG","tz1Zgo1nWvejNmjhzqEYPoRc57JhWrKhetAH","tz1gk9E28K1b6xQNw2wFLG9WERwzKDDwVNT9","tz1ftbRcC92jdKCTj91Qor9EkoJUcwqFFvH4","tz1hjorZSyLWMvupoBgWZsKbDFyca4B6wK36","tz1YNktyxUEzwpKq6QaWafka5CeKokJAEqi5","tz1KiuNcJNJ451967gXPwSDLvQmXryMvgTmL","tz1XDqZ4z2d5sRLkywYRwEMBQH5hHzaCaRre","tz1aQNn7eMC7kP1EyQetvocKS4TXmzEd5Gaw","tz1fnBJmAxoJbfAaMcCar3Aa8MFFH69dKYJv","tz1cZwxsasN1KorL8hYJPBy7gaBdNtsbjNvm","tz1b5qJrAgViKFPM5BdLeVasrxest7EWnf5z","tz1h87YTFnTCHKHTUEG2BxWq7DSRf861U8e6","tz1LPjWu1P11x1aizsD9wmqzey7ob66vvymg","tz1NBmAeLwsXxLcvAwJXM4nS199qkFcH4313","tz1cPL3YSMFiRC2yiMvSr7RLPVg5H5XhNvgW","tz1ZzGYssgyWGEZYVgCpMiYA1eBMWxRVrrqP","tz1ZgEQPeGYXeMLePL5sgR94RrRxwbsyK85z","tz1NcRjeevWE6E5B1pBsDexSB3cPJDTpQ9UL","tz1dBhoxSAU58V43Kvb7hG71iRxsfzmj7Z2y","tz1UwEDdQfpdcnYtfU7BRqXGPXbeaJmXk63c","tz1cpiGQS4rjtqnbndrbRnihtUjKH8aa1yHf","tz1eBdroBUqaUQ4h4xYXB26z67L8cj6ZyLsx","tz1Lk6U6erqH6EEuqMqk6u8gjqtJybZTirDs","tz1iAeHTDNZSwJjXaqbRyfyp96ych9j36tW1","tz1PDpbmvvBbFcEzctGfNt9mvvVTgTWxGGPZ","tz1RkXsWMgAtz85cD4iBckPC8xUFf3fPUFw2","tz1LcGrFcHLS9jgmWij4u7EWJDRrnU4xhVGL","tz1R752LZLPkJ1dhMs49zJqDcRZMic7nvL6g","tz1bpSiotWsQCz7jLjU4FV5K9PCZ79jinbs5","tz1bbsn8DMXJMEW7gVeuBbpMFBf65dq3VTtv","tz1V7aymCngroTZBxraVxp5kRtLh8qFRLcMN","tz1WxVo86ucoAE5RwCUe3tHTVV7r2GwRRvMN","tz1VXY3dXhE1s5huXnyFT7vfXgNroR5hTdjr","tz1LrYXWJ8Z3eKLSAYfCUaoZjpaqNrTqEd8w","tz1ZdYg6spcG9yk4ZXHHLcgNL2p8Y5wyhWYp","tz1PSqayN9mB1Wxk9SMjcGixX5PX9zRvtLZ8","tz1Pnke8YsfcDZfcqpZwpyFxLFCzEnSG9ews","tz1Qcgog1zYs6evgGP6aNEpGvDnCa7rAPvx5","tz1R5S8M9FAHs3eYjKY9ci3RD6iBgV8QCiLM","tz1iwGHsJu1Mq3eGGj2dmp5xj7N9kqxvuJzy","tz1TiEvTkhsuF5drB6Q9gDAQ3ejjp9GkMjdw","tz1LQPiugXCFuDit9232iqWGfj41PQGaCKvV","tz1TkAAz3CXdKiN4T8jAqxfWCBddQ39Vcewf","tz1ViczpurHDugWPgYzh5sE7Y7CDUB6qxseX","tz1NMf4uWqQNaVLFLjztKMpYphGdGSPu89pu","tz1ViNQrnj5aPjhZ9EQ1UuFcH6UUALqo3EUU","tz1Zqneq8GbnQ1G3g1Jnh9Vv4qGMnzLXbt3s","tz1dHH1xpCwZXE66WWK4SKDPLzF1JHX5KC96","tz1aXL6EouwJR3yMDK9jgD7zL5tpM12Vx3RY","tz1gtaaw2gzQjJmx8R17BRpYdkR9haE7i811","tz1aj98BzdTbicVFrqSYeAFjRwHGmRAwo2J6","tz1TS2ByHGPAgDb2AYv3bQWakG1tiH7uPwQb","tz1Q1t1xazVvEgN1dmxPCDXt8hqK9BY6qF5T","tz1fYzKUnmF873zvfcyUaX6yvRHbNTqakkSw","tz1g81md3cSqEtozkZggJnvRUNbMCbFHETqL","tz1bENLm1iQRnPsybc6HvrfR3W2nvuGErdFK","tz1TktVVrHESq3mdw6hpohXQFjbpGkdERsKT","tz1YSQ3PA5t3drYuAXYX5VE2mRcK143vjkKJ","tz1VFsJNozUp4mmtpv98bS2Q3B72qXWagDsq","tz1W47pyCkjpiWRAHmR9rHTmM5a1Yf9dtzkB","tz1Td2RL3zTSRZFac1SaXVxcXeAXHzTATsAU","KT1Tvqr6CTN3Lj2usqGPNG5Uog8JdgnDnhps","KT1CStwHYoGYzXENQamy2nc251enF5wk7Yjy","tz1UZuJ7Ae3m3rhsBjig955Fcy8vyfPYTDEr","tz1Qs9UZiWESjBJAA1SkqDm4eMDm78wo1QpN","tz1ZeovPMP8hy3PaJDg7AC8HP3rjNKpfRFes","tz1dkzioZTg4E3yVShiX1hdpAkE84wYhzRL8","KT1QhJRX8884Ccb9EKpNNFsbB27P6sRoqo4t","tz1Ww975VPEgcZdiTQGahABVkMf8yrcZb9A3","tz1RBv1smMW4xtTWG53fkzRZJGPmMotNJTEE","tz1cQD5tsef53FPzMmN5kLoLjvLD7VAGGQU5","tz1SUwahJKb5V9dxwk8N8quQWc5bQ1Zm6cva","tz1ciLo5K3g1pHPJKjEQp4WNvu6e9bxDtmwT","tz1ZTv3NVL4ry3qHZdmo7Y4rLjciHRcmptbA","tz1Z8TYYJsQgCqvvMtDo9Z56CzAtdVEi88NA","tz1hqbwdkxsh2MNaqoPVNycWA6cu1EKonKy8","tz1M4QhH1WQqrQWtsnDvyAqtfm3S7UcFAb78","tz1hxJteiSNWfmj12WdmHaPsdGtbZxZn1W9H","tz1cEnJ8Cu9AJcrfg9tajTmBEyTXPJM4teKP","tz1eSjS3Q2Rc36o2hHMRNX5pErtfGQRH9spH","tz1W7XN1Si3eGg4b6a7opsZ8EKjdiYVvmtfR","tz1gwyEjft79QrS5RKgLM3WozabfFeXn5kSC","tz1avdpHLEVbCKJPbFRD8pqySrZjPT64meHC","tz1NjziRqzmncVzU3CHXyo5BSVsn5CwXrUxS","tz1diyBGHXh6QR88Ps8guLmv8yCyyYTSoYjK","tz1RNVWP9MqtDSukfN9xZjNibKUMyq9ih2m1","tz1N1wThWmdxubf8y18ySS4KG69tR7icTHLn","tz1P96NympEyX7h4K5az2H8qr5vi7HyYGd7T","tz1Y7Ft4jEFvKE3dAHMUSwz9W3gXAWvN2ADR","tz1ND1DoSpPTtyooQQLxfNq7w8FR9dw9ZzuR","tz1SuBRZzrE98AeHRMy25zbNxwQfT5Wv1Cpp","tz1Q8HLHCFUFQVn38UeiSCNJUQsZxF2d6vtp","tz1TJrVZoozuTqzdeMakTnJQEogvm4q2531m","tz1TnfvKsMQajLzijg2EoKpE5yBGmXDiRQ1K","tz1gzg3UU6UMjVMGrcuetgnQ6btdVUGxYfyP","tz1KwcHHYUT7kMKeY6zXwUAq9Kuj3ryeX1iy","tz1WXuFmAiWqreKWY465FEps1Ei2XrAyHGGt","tz1TZafuJ3eWM3d3mckARNaP2oxBNyjLQVG6","tz1N7mPBeSrRsAgcukmRj2Twp6fvNshT1Uyc","tz1gFFSNHfQeNED9wYqThHMJZm5CNe6sFEUE","tz1L5xtPTd4XceH85n9W8vB3TjK45XGt63hC","tz1h6rqcMhAoEG2cMhiPjdJoemaHenY9mH7L","tz1bK9XE8khrUn7VxmEuNzF6FidGqpjTTXDs","tz1NUZw2MUG2AmSrtwTF3qdAUwgqBfnAGyu2","tz1ichwjUoGrUSawDKP6wMVLghJNX196aEBi","tz1UEEsxKrP1FH3RtUFt2esg7UdFQ74VeasS","tz1aNF6AiXu5uG6LENhYvJMwqYJc1bKTRuKr","KT1AG7uQwt3VX6JZah9KbD8n4sYSyaTUNqB3","tz1byac5DY5hA7dJ3uPvWvYAsWcAcdAGy953","tz1abWuJnBvT9o5TWzkEA2u3BbQ2mqFNoX3e","tz1Qin5ErcHtxtZV1njJ3cmzjw6GghWsEGGc"],"delegated_balance":"143547306933","deactivated":false,"grace_period":251}`

func createMockedAPI() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/chains/main/blocks/head/context/delegates/tz1V3yg82mcrPJbegqVCPn6bC8w1CSTRp3f8", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedTezosResponse); err != nil {
			panic(err)
		}
	})

	return r
}
