/**
 * Copyright 2002-2015 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textmapper.templates.types;

import java.io.IOException;
import java.io.Reader;
import java.text.MessageFormat;
import java.util.HashMap;
import java.util.Map;

public class TypesLexer {

	public static class Span {
		public Object value;
		public int symbol;
		public int state;
		public int line;
		public int offset;
		public int endoffset;
	}

	public interface Tokens {
		int Unavailable_ = -1;
		int eoi = 0;
		int identifier = 1;
		int scon = 2;
		int icon = 3;
		int bcon = 4;
		int _skip = 5;
		int DotDot = 6;
		int Dot = 7;
		int Mult = 8;
		int Semicolon = 9;
		int Comma = 10;
		int Colon = 11;
		int Assign = 12;
		int AssignGt = 13;
		int Lbrace = 14;
		int Rbrace = 15;
		int Lparen = 16;
		int Rparen = 17;
		int Lbrack = 18;
		int Rbrack = 19;
		int Lclass = 20;
		int Lextends = 21;
		int Lint = 22;
		int Lbool = 23;
		int Lstring = 24;
		int Lset = 25;
		int Lchoice = 26;
	}

	public interface ErrorReporter {
		void error(String message, int line, int offset, int endoffset);
	}

	public static final int TOKEN_SIZE = 2048;

	private Reader stream;
	final private ErrorReporter reporter;

	private CharSequence input;
	private int tokenOffset;
	private int l;
	private int charOffset;
	private int chr;

	private int state;

	private int tokenLine;
	private int currLine;
	private int currOffset;

	private String unescape(String s, int start, int end) {
		StringBuilder sb = new StringBuilder();
		end = Math.min(end, s.length());
		for(int i = start; i < end; i++) {
			char c = s.charAt(i);
			if(c == '\\') {
				if(++i == end) {
					break;
				}
				c = s.charAt(i);
				if(c == 'u' || c == 'x') {
					// FIXME process unicode
				} else if(c == 'n') {
					sb.append('\n');
				} else if(c == 'r') {
					sb.append('\r');
				} else if(c == 't') {
					sb.append('\t');
				} else {
					sb.append(c);
				}
			} else {
				sb.append(c);
			}
		} 
		return sb.toString();
	}

	public TypesLexer(CharSequence input, ErrorReporter reporter) throws IOException {
		this.reporter = reporter;
		reset(input);
	}

	public void reset(CharSequence input) throws IOException {
		this.state = 0;
		tokenLine = currLine = 1;
		currOffset = 0;
		this.input = input;
		tokenOffset = l = 0;
		charOffset = l;
		chr = l < input.length() ? input.charAt(l++) : -1;
		if (chr >= Character.MIN_HIGH_SURROGATE && chr <= Character.MAX_HIGH_SURROGATE && l < input.length() &&
				Character.isLowSurrogate(input.charAt(l))) {
			chr = Character.toCodePoint((char) chr, input.charAt(l++));
		}
	}

	protected void advance() {
		if (chr == -1) return;
		currOffset += l - charOffset;
		if (chr == '\n') {
			currLine++;
		}
		charOffset = l;
		chr = l < input.length() ? input.charAt(l++) : -1;
		if (chr >= Character.MIN_HIGH_SURROGATE && chr <= Character.MAX_HIGH_SURROGATE && l < input.length() &&
				Character.isLowSurrogate(input.charAt(l))) {
			chr = Character.toCodePoint((char) chr, input.charAt(l++));
		}
	}

	public int getState() {
		return state;
	}

	public void setState(int state) {
		this.state = state;
	}

	public int getTokenLine() {
		return tokenLine;
	}

	public int getLine() {
		return currLine;
	}

	public void setLine(int currLine) {
		this.currLine = currLine;
	}

	public int getOffset() {
		return currOffset;
	}

	public void setOffset(int currOffset) {
		this.currOffset = currOffset;
	}

	public String tokenText() {
		return input.subSequence(tokenOffset, charOffset).toString();
	}

	public int tokenSize() {
		return charOffset - tokenOffset;
	}

	private static final short tmCharClass[] = {
		1, 1, 1, 1, 1, 1, 1, 1, 1, 31, 4, 1, 1, 31, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		31, 1, 5, 15, 1, 1, 1, 2, 25, 26, 17, 1, 19, 6, 16, 1,
		30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 20, 18, 1, 21, 22, 1,
		1, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
		29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 27, 3, 28, 1, 29,
		1, 12, 29, 29, 29, 10, 11, 29, 29, 29, 29, 29, 13, 29, 29, 29,
		29, 29, 8, 14, 7, 9, 29, 29, 29, 29, 29, 23, 1, 24, 1, 1
	};

	private static final int[] tmRuleSymbol = unpack_int(27,
		"\1\0\2\0\3\0\4\0\5\0\5\0\6\0\7\0\10\0\11\0\12\0\13\0\14\0\15\0\16\0\17\0\20\0\21" +
		"\0\22\0\23\0\24\0\25\0\26\0\27\0\30\0\31\0\32\0");

	private static final int tmClassesCount = 32;

	private static final short[] tmGoto = unpack_vc_short(1088,
		"\1\ufffe\1\uffff\1\1\1\uffff\1\2\1\3\1\4\1\5\3\6\1\7\3\6\1\10\1\11\1\12\1\13\1\14" +
		"\1\15\1\16\1\uffff\1\17\1\20\1\21\1\22\1\23\1\24\1\6\1\25\1\2\1\uffff\1\1\1\26\1" +
		"\27\1\uffff\33\1\4\ufff9\1\2\32\ufff9\1\2\1\uffff\2\3\1\30\1\uffff\1\31\32\3\36\uffff" +
		"\1\25\1\uffff\7\ufffd\1\6\1\32\6\6\16\ufffd\2\6\10\ufffd\10\6\16\ufffd\2\6\10\ufffd" +
		"\5\6\1\33\2\6\16\ufffd\2\6\1\ufffd\1\ufff8\3\10\1\ufff8\33\10\20\ufff6\1\34\17\ufff6" +
		"\40\ufff5\40\ufff4\40\ufff3\40\ufff2\26\ufff1\1\35\11\ufff1\40\uffef\40\uffee\40" +
		"\uffed\40\uffec\40\uffeb\40\uffea\36\ufffb\1\25\1\ufffb\40\ufffd\1\uffff\3\1\1\uffff" +
		"\33\1\1\uffff\3\3\1\uffff\33\3\40\ufffc\7\ufffd\2\6\1\36\5\6\16\ufffd\2\6\10\ufffd" +
		"\6\6\1\37\1\6\16\ufffd\2\6\1\ufffd\40\ufff7\40\ufff0\7\ufffd\3\6\1\40\4\6\16\ufffd" +
		"\2\6\10\ufffd\7\6\1\41\16\ufffd\2\6\1\ufffd\7\ufffa\10\6\16\ufffa\2\6\1\ufffa\7\ufffd" +
		"\3\6\1\40\4\6\16\ufffd\2\6\1\ufffd");

	private static short[] unpack_vc_short(int size, String... st) {
		short[] res = new short[size];
		int t = 0;
		int count = 0;
		for (String s : st) {
			int slen = s.length();
			for (int i = 0; i < slen; ) {
				count = i > 0 || count == 0 ? s.charAt(i++) : count;
				if (i < slen) {
					short val = (short) s.charAt(i++);
					while (count-- > 0) res[t++] = val;
				}
			}
		}
		assert res.length == t;
		return res;
	}

	private static int mapCharacter(int chr) {
		if (chr >= 0 && chr < 128) return tmCharClass[chr];
		return chr == -1 ? 0 : 1;
	}

	public Span next() throws IOException {
		Span token = new Span();
		int state;

		tokenloop:
		do {
			token.offset = currOffset;
			tokenLine = token.line = currLine;
			tokenOffset = charOffset;

			for (state = this.state; state >= 0; ) {
				state = tmGoto[state * tmClassesCount + mapCharacter(chr)];
				if (state == -1 && chr == -1) {
					token.endoffset = currOffset;
					token.symbol = 0;
					token.value = null;
					reporter.error("Unexpected end of input reached", token.line, token.offset, token.endoffset);
					token.offset = currOffset;
					break tokenloop;
				}
				if (state >= -1 && chr != -1) {
					currOffset += l - charOffset;
					if (chr == '\n') {
						currLine++;
					}
					charOffset = l;
					chr = l < input.length() ? input.charAt(l++) : -1;
					if (chr >= Character.MIN_HIGH_SURROGATE && chr <= Character.MAX_HIGH_SURROGATE && l < input.length() &&
							Character.isLowSurrogate(input.charAt(l))) {
						chr = Character.toCodePoint((char) chr, input.charAt(l++));
					}
				}
			}
			token.endoffset = currOffset;

			if (state == -1) {
				reporter.error(MessageFormat.format("invalid lexeme at line {0}: `{1}`, skipped", currLine, tokenText()), token.line, token.offset, token.endoffset);
				token.symbol = -1;
				continue;
			}

			if (state == -2) {
				token.symbol = Tokens.eoi;
				token.value = null;
				break tokenloop;
			}

			token.symbol = tmRuleSymbol[-state - 3];
			token.value = null;

		} while (token.symbol == -1 || !createToken(token, -state - 3));
		return token;
	}

	protected int charAt(int i) {
		if (i == 0) return chr;
		i += l - 1;
		int res = i < input.length() ? input.charAt(i++) : -1;
		if (res >= Character.MIN_HIGH_SURROGATE && res <= Character.MAX_HIGH_SURROGATE && i < input.length() &&
				Character.isLowSurrogate(input.charAt(i))) {
			res = Character.toCodePoint((char) res, input.charAt(i++));
		}
		return res;
	}

	protected boolean createToken(Span token, int ruleIndex) throws IOException {
		boolean spaceToken = false;
		switch (ruleIndex) {
			case 0:
				return createIdentifierToken(token, ruleIndex);
			case 1: // scon: /"([^\n\\"]|\\.)*"/
				{ token.value = unescape(tokenText(), 1, tokenSize()-1); }
				break;
			case 2: // icon: /\-?[0-9]+/
				{ token.value = Integer.parseInt(tokenText()); }
				break;
			case 3: // bcon: /true|false/
				{ token.value = tokenText().equals("true"); }
				break;
			case 4: // _skip: /[\n\t\r ]+/
				spaceToken = true;
				break;
			case 5: // _skip: /#.*/
				spaceToken = true;
				break;
		}
		return !(spaceToken);
	}

	private static Map<String,Integer> subTokensOfIdentifier = new HashMap<>();
	static {
		subTokensOfIdentifier.put("class", 20);
		subTokensOfIdentifier.put("extends", 21);
		subTokensOfIdentifier.put("int", 22);
		subTokensOfIdentifier.put("bool", 23);
		subTokensOfIdentifier.put("string", 24);
		subTokensOfIdentifier.put("set", 25);
		subTokensOfIdentifier.put("choice", 26);
	}

	protected boolean createIdentifierToken(Span token, int ruleIndex) {
		Integer replacement = subTokensOfIdentifier.get(tokenText());
		if (replacement != null) {
			ruleIndex = replacement;
			token.symbol = tmRuleSymbol[ruleIndex];
		}
		boolean spaceToken = false;
		switch(ruleIndex) {
			case 0:	// <default>
				{ token.value = tokenText(); }
				break;
		}
		return !(spaceToken);
	}

	/* package */ static int[] unpack_int(int size, String... st) {
		int[] res = new int[size];
		boolean second = false;
		char first = 0;
		int t = 0;
		for (String s : st) {
			int slen = s.length();
			for (int i = 0; i < slen; i++) {
				if (second) {
					res[t++] = (s.charAt(i) << 16) + first;
				} else {
					first = s.charAt(i);
				}
				second = !second;
			}
		}
		assert !second;
		assert res.length == t;
		return res;
	}

}
