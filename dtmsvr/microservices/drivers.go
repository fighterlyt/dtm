package microservices

import (
	// load the microserver drivers
	_ "github.com/fighterlyt/dtmdriver-dapr"
	_ "github.com/fighterlyt/dtmdriver-ego"
	_ "github.com/fighterlyt/dtmdriver-gozero"
	_ "github.com/fighterlyt/dtmdriver-kratos"
	_ "github.com/fighterlyt/dtmdriver-polaris"
	_ "github.com/fighterlyt/dtmdriver-springcloud"
	_ "github.com/zhufuyi/dtmdriver-sponge"
)
