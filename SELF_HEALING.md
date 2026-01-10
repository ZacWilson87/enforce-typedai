# Self-Healing (Repair) Logic

## Purpose
Define how invalid model output may be repaired in a controlled, deterministic manner.

Self-healing exists to recover from *structural* failures, not semantic ambiguity.

---

## When Repair Is Triggered

Repair is triggered **only** when:
- Model output fails schema validation
- The failure is structural (e.g. missing fields, wrong types)
- Repair is explicitly enabled in the request

Repair must NOT trigger for:
- Semantic errors
- Business logic violations
- Ambiguous or underspecified output

---

## Repair Flow

```
Model Output
  ↓
Schema Validation Failure
  ↓
Repair Attempt (bounded)
  ↓
Re-validation
  ├─ Success → Return Result (Repaired = true)
  └─ Failure → Retry or Exit
```

---

## Repair Rules

- Repair attempts are **bounded**
- Default max repair attempts: 1
- Repair attempts must not mutate:
  - Original schema
  - Original request intent
- Each repair attempt must be logged and counted
- Repair must terminate deterministically

No infinite loops.
No adaptive retries.
No hidden state.

---

## Repair Prompt Constraints

Repair prompts must:
- Include the validation error
- Include the original schema
- Include the invalid output verbatim
- Instruct the model to output **only valid structured data**

Repair prompts must NOT:
- Add new fields
- Infer missing intent
- Introduce creativity
- Change output meaning

---

## Output Contract

If repair succeeds:
- Result must indicate `Repaired = true`
- Attempt count must include repair attempts

If repair fails:
- Return a deterministic repair error
- Include last invalid output
- Do not retry beyond configured limits

---

## Failure Modes

- RepairDisabled
- RepairExhausted
- RepairInvalidOutput

All failure modes must be explicit and typed.

---

## Notes for AI

- Treat repair as exception handling, not default flow
- Never introduce randomness or backoff
- Do not escalate repair into agent-like behavior
- If unsure, fail explicitly
