// generated by Textmapper; DO NOT EDIT

#ifndef EXAMPLES_JSON_PARSER_H_
#define EXAMPLES_JSON_PARSER_H_

#include <array>
#include <cstdint>
#include <ostream>
#include <string>
#include <utility>
#include <vector>

#include "absl/base/attributes.h"
#include "absl/functional/function_ref.h"
#include "absl/log/log.h"
#include "absl/status/status.h"
#include "absl/strings/str_format.h"
#include "lexer.h"

namespace json {

struct symbol {
  int32_t symbol = 0;
  int64_t offset = 0;
  int64_t endoffset = 0;
};

struct stackEntry {
  symbol sym;
  int8_t state = 0;
};

enum class NodeType {
  NoType,
  EmptyObject,
  JSONArray,
  JSONMember,
  JSONObject,
  JSONText,
  JSONValue,
  MultiLineComment,
  InvalidToken,
  JsonString,
  NonExistingType,
  NodeTypeMax
};

constexpr inline std::array<absl::string_view,
                            static_cast<size_t>(NodeType::NodeTypeMax)>
    nodeTypeStr = {
        "NONE",         "EmptyObject", "JSONArray",       "JSONMember",
        "JSONObject",   "JSONText",    "JSONValue",       "MultiLineComment",
        "InvalidToken", "JsonString",  "NonExistingType",
};

inline std::ostream& operator<<(std::ostream& os, NodeType t) {
  int i = static_cast<int>(t);
  if (i >= 0 && i < nodeTypeStr.size()) {
    return os << nodeTypeStr[i];
  }
  return os << "node(" << i << ")";
}

constexpr inline bool debugSyntax = false;
constexpr inline int startStackSize = 256;
constexpr inline int startTokenBufferSize = 16;
constexpr inline int32_t noToken = static_cast<int32_t>(Token::UNAVAILABLE);
constexpr inline int32_t eoiToken = static_cast<int32_t>(Token::EOI);

ABSL_MUST_USE_RESULT std::string symbolName(int32_t sym);

class Parser {
 public:
  template <typename Listener>
  explicit Parser(Listener&& listener)
      : listener_(std::forward<Listener>(listener)) {
    pending_symbols_.reserve(startTokenBufferSize);
  }

  absl::Status Parse(Lexer& lexer) { return Parse(1, 44, lexer); }

 private:
  void reportIgnoredToken(symbol sym);
  void fetchNext(Lexer& lexer, std::vector<stackEntry>& stack);
  absl::Status applyRule(int32_t rule, stackEntry& lhs,
                         [[maybe_unused]] absl::Span<const stackEntry> rhs,
                         Lexer& lexer);
  absl::Status Parse(int8_t start, int8_t end, Lexer& lexer);

  symbol next_symbol_;
  // Tokens to be reported with the next shift. Only non-empty when next.symbol
  // != noToken.
  std::vector<symbol> pending_symbols_;
  absl::FunctionRef<void(NodeType, int64_t, int64_t)> listener_;

  friend std::ostream& operator<<(std::ostream& os, const Parser& parser) {
    return os << "JSONParser next " << symbolName(parser.next_symbol_.symbol)
              << " and pending " << parser.pending_symbols_.size()
              << " symbols";
  }
};

}  // namespace json

#endif  // EXAMPLES_JSON_PARSER_H_
