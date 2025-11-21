

type Policy {
    enum Effect {
        unspecified  = 0;
        allow = 1;
        deny  = 2;
    }

    // ["<", "x", "12"]
    // ["match", "user", "abc*"]
    type Condition {
        operator: String @1
        key: String @2
        values: [Value] @3
    }

    id: String @1
    name: String @2
    description: String @3

    domain: String @4

    effect: Effect @10
    subjects: [Subject] @11 @db.json
    permissions: [Permission] @12 @db.json
    conditions: [Condition] @13 @db.json

    create_time: Timestamp @1000
    update_time: Timestamp @1001
}

